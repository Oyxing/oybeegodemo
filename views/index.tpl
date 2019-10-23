<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  

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
   
     <form action="/api/getusermag" method="post">
    <table>
        <tbody>
        <tr> 
            <td>用户名</td>
            <td>
                <input type="text" name ="name" value="">
            </td>
        </tr>
        <tr>
            <td>密码</td>
            <td>
                <input type="password" name ="pwd" value="">
            </td>
 
        </tr>
        <tr>
            <td>年龄</td>
            <td>
                <input type="text" name ="age" value="">
            </td>
 
        </tr>
        <tr>
            <td>bool</td>
            <td>
                <input type="text" name ="bool" value="">
            </td>
 
        </tr>
        <tr>
            <td>
                <input type="submit" value="确认">
            </td>
            <td>
                <input type="reset" value="重置">
            </td>
 
        </tr>
        </tbody>
    </table>
</form>

<h2>
{{.Username}}
</h2>
 <h1 class="logo">Welcome to Beego</h1>
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
		<h1>{{.sessus}}</h1>
		
		{{range .data}}
	        {{.}}
		
	    {{end}}
    </div>
  </footer>
  	<div class="backdrop">
	
	</div>

  <script src="/static/js/reload.min.js"></script>
  <script>
    function buttons(){
      console.log({{.Website}})
      console.log({{.Clickfun}})
    }
  </script>
</body>
</html>
