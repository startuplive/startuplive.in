package views

import (
	. "github.com/ungerik/go-start/view"
)

var RenderUserVoice = Render(
	func(ctx *Context) (err error) {

		// if !Config.IsProductionServer {
		// 	return nil
		// }

		ctx.Response.Write([]byte(`<script type="text/javascript">
  var uvOptions = {};
  (function() {
    var uv = document.createElement('script'); uv.type = 'text/javascript'; uv.async = true;
    uv.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'widget.uservoice.com/8UZUw5obSbwMoGONxybPyg.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(uv, s);
  })();
</script>`))

		return nil
	},
)
