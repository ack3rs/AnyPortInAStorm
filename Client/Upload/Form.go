package Upload

import (
	"net/http"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {

	html := []byte(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Any Port In a Storm</title>
  </head>
  <body>

	<h1>Any Port In a Storm</h1>

    <form enctype="multipart/form-data" action="/upload" method="post">
      <input type="file" name="portsFile" />
      <input type="submit" value="Upload" />
    </form>
  </body>
</html>`)

	w.Write(html)

}
