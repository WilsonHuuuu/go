<h1>golang标准库ioutil包</h1>
<table>
    <tr>
        <th>名称</th>
        <th>作用</th>
    </tr>
    <tr>
        <td>ReadAll</td>        
        <td>读取数据，返回读到的字节切片</td>
    </tr>
    <tr>
        <td>ReadDir</td>        
        <td>读取一个目录，返回目录入口数组[]os.FileInfo</td>
    </tr>
    <tr>
        <td>ReadFile</td>        
        <td>读一个文件，返回文件内容（字节切片）</td>
    </tr>
    <tr>
        <td>WriteFile</td>        
        <td>根据文件路径，写入字节切片</td>
    </tr>
    <tr>
        <td>TempDir</td>        
        <td>在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径</td>
    </tr>
    <tr>
        <td>TempFile</td>        
        <td>在一个目录中创建指定前缀名的临时文件，返回os.File</td>
    </tr>
</table>