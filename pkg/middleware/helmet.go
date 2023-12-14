package middleware

import "github.com/gofiber/fiber/v2/middleware/helmet"

var HelmetConfig = helmet.Config{
	XSSProtection:             "1",
	ContentTypeNosniff:        "nosniff",
	XFrameOptions:             "SAMEORIGIN",
	HSTSExcludeSubdomains:     false,
	ContentSecurityPolicy:     "default-src 'self';base-uri 'self';font-src 'self' https: data:;form-action 'self';frame-ancestors 'self';img-src 'self' data:;object-src 'none';script-src 'self';script-src-attr 'none';style-src 'self' https: 'unsafe-inline';upgrade-insecure-requests",
	CSPReportOnly:             false,
	HSTSPreloadEnabled:        true,
	ReferrerPolicy:            "no-referrer",
	PermissionPolicy:          "geolocation=(self)",
	CrossOriginEmbedderPolicy: "require-corp",
	CrossOriginOpenerPolicy:   "same-origin",
	CrossOriginResourcePolicy: "same-origin",
	OriginAgentCluster:        "?1",
	XDNSPrefetchControl:       "off",
	XDownloadOptions:          "noopen",
	XPermittedCrossDomain:     "none",
}
