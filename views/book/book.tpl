<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>后台登录</title>
     <script src="//code.jquery.com/jquery-1.11.3.min.js"></script>
    <style>

        .main {
            margin: 0 auto;
            width: 400px;
            border: 1px solid;
            border-color: #eeeeee;
            border-radius: 5px;
            margin-top: 100px;
        }
    </style>

</head>
<body>
    <embed   width="800" height="1500" quality="high"  src="/static/books/{{.Name}}" align=center />

</body>
<script>
    $(function () {
    // init...
    setSize()
});
</script>
<script>
    function setSize() {
    $('embed').height(window.width * 0.7);
};
</script>
</html>