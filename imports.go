package crowdsec

import (
	// Import the default CrowdSec modules. Primary reason this
	// file exists is to satisfy the Caddy documentation and download
	// pages to list the modules correctly.
	_ "github.com/fhuteau/caddy-crowdsec-bouncer/appsec"
	_ "github.com/fhuteau/caddy-crowdsec-bouncer/crowdsec"
	_ "github.com/fhuteau/caddy-crowdsec-bouncer/http"
	_ "github.com/fhuteau/caddy-crowdsec-bouncer/layer4"
)
