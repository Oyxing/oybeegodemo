<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script src="/static/js/jquery.js"></script>
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 250px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
	.backdrop{
		color:red;
	
		background-color:#ccc;
	}
  </style>
</head>

<body>
  <header>
      <h2>
      {{.Username}}
      </h2>
      <h1 class="logo">Welcome to Beego</h1>
            <img style="width:280px;height:280px;" id="billImage" src="data:image/png;base64,{{.Image}}" />
          {{.data}}
          <div class="description">
            Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
          </div>
      
          <button onclick="buttons()">sadas</button>
        </header>
      <div>
    </div>
    <footer>
      <div class="author">
        Official website:
        <a href="http://{{.Website}}">{{.Website}}</a> /
        Contact me:
        <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>

      </div>
    </footer>
  	<div class="backdrop">
	</div>
  <script src="/static/js/reload.min.js"></script>
  <script>
      $.ajax({
          url: "/static/js/industryclassify.json",
          type: "GET",
          success: function(res) {
            console.log(res)
          }
      })
  </script>
</body>
</html>
