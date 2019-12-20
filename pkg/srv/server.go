package srv

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"k8s.io/klog"
)

// Server is the primary struct upon which all of the server capabilities are built
type Server struct {
	HTTP *http.Server

	cfg *viper.Viper
}

// New generates a new instance of Server
//
// Arguments:
//     cfg (*viper.Viper): A viper instance to configure the server with
//
// Returns:
//     (*Server): A pointer to the new instance of Server, nil if an error occurred
//     (error):   An error if one exists, nil otherwise
func New(cfg *viper.Viper) (*Server, error) {
	var S Server
	S.cfg = cfg

	idleTO, err := S.parseTimeout(S.cfg.GetString("http.idle_timeout"))
	if err != nil {
		return nil, err
	}
	readTO, err := S.parseTimeout(S.cfg.GetString("http.read_timeout"))
	if err != nil {
		return nil, err
	}
	writeTO, err := S.parseTimeout(S.cfg.GetString("http.write_timeout"))
	if err != nil {
		return nil, err
	}

	if S.cfg.GetBool("http.tls.enable") {
		return nil, fmt.Errorf("tls is unsupported at the moment")
	} else {
		S.HTTP = &http.Server{
			Addr:         fmt.Sprintf("%s:%d", S.cfg.GetString("http.bind_address"), cfg.GetInt("http.bind_port")),
			IdleTimeout:  idleTO,
			ReadTimeout:  readTO,
			WriteTimeout: writeTO,
			Handler:      S.newRouter(),
		}
	}

	return &S, nil
}

// Run will run the HTTP(S) server
//
// Arguments:
//     None
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (S *Server) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	runChan := make(chan os.Signal, 1)

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	if S.cfg.GetBool("http.tls.enable") {
		klog.Infof("running the HTTPS server on %s", S.HTTP.Addr)
		go func() {
			if err := S.HTTP.ListenAndServeTLS("", ""); err != nil {
				klog.Fatal(err)
			}
		}()
	} else {
		klog.Infof("running the HTTP server on %s", S.HTTP.Addr)
		go func() {
			if err := S.HTTP.ListenAndServe(); err != nil {
				klog.Fatal(err)
			}
		}()
	}

	<-runChan

	klog.Info("shutting down the server due to interrupt")
	if err := S.HTTP.Shutdown(ctx); err != nil {
		klog.Fatal(err)
	}
}

func (S *Server) parseTimeout(s string) (time.Duration, error) {
	str := strings.ToLower(s)
	switch {
	case strings.HasSuffix(str, "μ"):
		i, err := strconv.Atoi(strings.TrimSuffix(str, "μ"))
		if err != nil {
			return 0, err
		}
		return time.Duration(i) * time.Microsecond, nil
	case strings.HasSuffix(str, "ms"):
		i, err := strconv.Atoi(strings.TrimSuffix(str, "ms"))
		if err != nil {
			return 0, err
		}
		return time.Duration(i) * time.Millisecond, nil
	case strings.HasSuffix(str, "s"):
		i, err := strconv.Atoi(strings.TrimSuffix(str, "s"))
		if err != nil {
			return 0, err
		}
		return time.Duration(i) * time.Second, nil
	case strings.HasSuffix(str, "m"):
		i, err := strconv.Atoi(strings.TrimSuffix(str, "m"))
		if err != nil {
			return 0, err
		}
		return time.Duration(i) * time.Minute, nil
	case strings.HasSuffix(str, "h"):
		i, err := strconv.Atoi(strings.TrimSuffix(str, "h"))
		if err != nil {
			return 0, err
		}
		return time.Duration(i) * time.Hour, nil
	default:
		return 0, fmt.Errorf("'%s' contains in invalid time suffix, must be one of [μ|ms|s|m|h]", s)
	}
}
