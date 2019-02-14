
layui.use(['form', 'table', 'jquery'], function(){
    var form = layui.form
    ,table = layui.table
    ,$ = layui.jquery;

     //监听提交
  form.on('submit(formDemo)', function(data){
      table.reload('table', {
        where: data.field //设定异步数据接口的额外参数
      });
      return false;
  });

  form.on('submit(save)', function(data){
    var method = "POST"
    
    if (data.field.id > 0) {
        method = "PUT"
    } 
    
     $.ajax({
         url: '',
         type: 'json',
         method: method,
         data: data.field,
         success: function (ret) {
            if (ret.code == 0) {
                layer.msg("操作成功", {icon: 1, time: 1000}, function () {
                    parent.layer.closeAll()
                    window.parent.location.reload()
                })
            } else {
                layer.msg(ret.msg, {icon: 2})
            }
         }
     })
      return false;
  });


  table.on("toolbar(table)", function (obj) {
      var layEevent = obj.event;
      var mid = $('#add').data('mid');
      
      if (layEevent == "add") {
        layer.open({
            type: 2, 
            title: '添加权限',
            area: ['600px', '500px'],
            content: '/privileges/' + mid //这里content是一个URL，如果你不想让iframe出现滚动条，你还可以content: ['http://sentsin.com', 'no']
          }); 
      }
  })

  table.on("tool(table)", function (obj) {
        var data = obj.data;
        var layEevent = obj.event;
        if (layEevent == 'edit') {
            layer.open({
                type: 2, 
                title: '编辑权限',
                area: ['600px', '500px'],
                content: '/privileges/' + data.Id //这里content是一个URL，如果你不想让iframe出现滚动条，你还可以content: ['http://sentsin.com', 'no']
              }); 
        } else if (layEevent == 'del') {
            layer.confirm('is not?', {icon: 3, title:'提示'}, function(index){
                $.ajax({
                    url: '/privileges/' + data.Id,
                    type: 'json',
                    method: 'Delete',
                    success: function (ret) {
                       if (ret.code == 0) {
                           layer.msg("操作成功", {icon: 1, time: 1000}, function () {
                            layer.close(index);
                             window.location.reload()
                           })
                       } else {
                           layer.msg(ret.msg, {icon: 2})
                       }
                    }
                })
                
              });
           
        }
  })
})