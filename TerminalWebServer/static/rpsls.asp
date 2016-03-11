<body>
<%
Dim name, age
name = Request.QueryString("input")
age = Request.QueryString("input")
Response.Write("Name: " & name & "<br />")
Response.Write("Age: " & age & "<br />")
%>
</body>