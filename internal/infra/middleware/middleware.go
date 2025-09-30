package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// SetupMiddleware 设置中间件
func SetupMiddleware(e *echo.Echo, logger *zap.Logger) {
	// CORS 中间件
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// 请求恢复中间件
	e.Use(middleware.Recover())

	// 请求ID中间件
	e.Use(middleware.RequestID())

	// 超时中间件
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	// 限流中间件
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// 请求日志中间件
	e.Use(RequestLogger(logger))

	// 安全头中间件
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	// 压缩中间件
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
}

// RequestLogger 请求日志中间件
func RequestLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("HTTP Request",
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("user_agent", v.UserAgent),
				zap.String("request_id", v.RequestID),
			)
			return nil
		},
	})
}

// ErrorHandler 自定义错误处理器
func ErrorHandler(logger *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		code := 500
		message := "Internal Server Error"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			if msg, ok := he.Message.(string); ok {
				message = msg
			}
		}

		// 记录错误日志
		logger.Error("HTTP Error",
			zap.Error(err),
			zap.Int("status", code),
			zap.String("method", c.Request().Method),
			zap.String("uri", c.Request().RequestURI),
			zap.String("remote_ip", c.RealIP()),
		)

		// 发送错误响应
		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				err = c.NoContent(code)
			} else {
				err = c.JSON(code, map[string]interface{}{
					"error":     message,
					"code":      "HTTP_ERROR",
					"status":    code,
					"path":      c.Request().RequestURI,
					"method":    c.Request().Method,
					"timestamp": time.Now().Unix(),
				})
			}
			if err != nil {
				logger.Error("Failed to send error response", zap.Error(err))
			}
		}
	}
}
