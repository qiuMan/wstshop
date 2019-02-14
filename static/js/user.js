
layui.use(['form', 'table'], function(){
    var form = layui.form
    ,table = layui.table;
     //监听提交
  form.on('submit(search)', function(data){
      table.reload('table', {
        where: data.field //设定异步数据接口的额外参数
      });
      return false;
  });
})