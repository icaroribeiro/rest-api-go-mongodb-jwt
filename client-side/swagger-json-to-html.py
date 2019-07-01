#!/usr/bin/python
"""
Usage: python swagger-json-to-html.py < openapi.json > index.html
"""
import json, sys

TEMPLATE = """
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>RESTful API</title>
    <link rel="stylesheet" type="text/css" href="./swagger-ui/3.23.0/swagger-ui.css" >
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>
    <script src="./swagger-ui/3.23.0/swagger-ui-bundle.js"> </script>
    <script src="./swagger-ui/3.23.0/swagger-ui-standalone-preset.js"> </script>
    <script src="./swagger-ui/3.23.0/swagger-ui.js"> </script>
    <script>
    window.onload = function() {

      var spec = %s;

      const HideTopbarPlugin = function() {
        return {
          components: {
            Topbar: function() {
              return null
            }
          }
        }
      }

      const DisableTryItOutPlugin = function() {
        return {
          statePlugins: {
            spec: {
              wrapSelectors: {
                allowTryItOutFor: () => () => false
              }
            }
          }
        }
      }

      const ui = SwaggerUIBundle({
        spec: spec,
        dom_id: '#swagger-ui',
        <!-- withCredentials: true, -->
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl,
          HideTopbarPlugin
        ],
        layout: "StandaloneLayout"
      })

      window.ui = ui
    }
  </script>
  </body>
</html>
"""

spec = json.load(sys.stdin)
sys.stdout.write(TEMPLATE % json.dumps(spec))
