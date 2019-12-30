package mail

const (
	head = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta content="telephone=no" name="format-detection">
    <title></title>
    <style type="text/css">
        /* Basics */
        html, body {
            margin: 0 !important;
            padding: 0 !important;
            background-color: #ffffff;
			height: 100% !important;
			width: 100% !important;
			-webkit-text-size-adjust: 100%;
            -ms-text-size-adjust: 100%;
		}
        table {
            font-family: sans-serif;
            color: #333333;
			border-spacing: 0 !important;
    		border-collapse: collapse !important;
    		margin: 0 auto !important;
        }
        td {
            padding: 0;
        }
        img {
            border: 0;
        }
        div[style*="margin: 16px 0"] {
            margin:0 !important;
        }
        .center {
            text-align: center;
        }
        .wrapper {
            width: 100%;
            table-layout: fixed;
        }
        .webkit {
            max-width: 600px;
            margin: 0 auto;
        }
        .outer {
            Margin: 0 auto;
            width: 100%;
            max-width: 600px;
        }
		span {
			font-size: 14px;
		}
        .inner {
            padding-left: 10px;
			padding-right: 10px;
        }
        .contents {
            width: 100%;
        }
        p {
            Margin: 0;
        }
        a {
            color: #ee6a56;
            text-decoration: underline;
        }
		.db {
            display: block;
		},
		.di {
            display: inline;
        }
		.m0 {
			margin: 0 !important;
		},
		.mr4  {  
			margin-right: 90px;
		}
        .h1 {
            font-size: 21px;
            font-weight: bold;
            Margin-bottom: 18px;
        }
        .h2 {
            font-size: 18px;
            font-weight: bold;
            Margin-bottom: 12px;
        }
        .full-width-image img {
            width: 100%;
            max-width: 600px;
            height: auto;
        }
        .one-column p {
            font-size: 14px;
			margin-bottom: 10px;
        }
		.one-column i {
            font-size: 14px;
			margin-bottom: 10px;
        }
		.tc  { 
			text-align: center;
		}
		.tr  { 
			text-align: right; 
		}
		.tl  { 
			text-align: left; 
		}
		.fl  {
			float: left;
		}
		.fr  {
			float: right;
		}
    </style>
</head>
<body>
<center class="wrapper">
    <div class="webkit">
        <table class="outer" align="center">`

	footer = `</table>
    </div>
</center>
</body>
</html>`
)