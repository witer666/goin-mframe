<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Gin Form</title>
</head>
<body>
<form name="gfrom" action="/admin/post" method="post" enctype="application/x-www-form-urlencoded">
  <div>姓名：<input type="text" name="uname"><div>
  <div>性别：<input name="sex" type="radio" value="male"><input name="sex" type="radio" value="female"><div>
  <div>爱好：<input type="checkbox" name="likes" value="旅游" />旅游  
    <input type="checkbox" name="likes" value="跑步" />跑步
    <input type="checkbox" name="likes" value="看书" />看书<div>
  <div>省份：<select name="province">
      <option value="河北">河北</option>
      <option value="陕西">陕西</option>
      <option value="北京">北京</option>
      <option value="天津">天津</option>
  </select>
  <div>
  <input type="submit" value="提交">
</form>
</body>
</html>