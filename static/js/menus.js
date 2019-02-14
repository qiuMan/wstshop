var editObj=null,ptable=null,treeGrid=null,tableId='treeTable',layer=null;

layui.config({
    base: '/static/layui/extend/'
}).extend({
    treeGrid:'treeGrid'
}).use(['jquery','treeGrid','layer'], function(){
    var $=layui.jquery;
    treeGrid = layui.treeGrid;//很重要
    layer=layui.layer;
    ptable=treeGrid.render({
        id:tableId
        ,elem: '#'+tableId
        ,idField:'id'
        ,url:'/menus/index'
        ,cellMinWidth: 100
        ,treeId:'MenuId'//树形id字段名称
        ,treeUpId:'ParentId'//树形父id字段名称
        ,treeShowName:'MenuName'//以树形式显示的字段
        ,cols: [[
            {field:'MenuId',width:100, title: 'MenuId'}
            ,{field:'ParentId',width:100, title: 'ParentId'}
            ,{field:'MenuName', edit:'text', title: '名称'}
            ,{width:200,title: '操作', align:'left'/*toolbar: '#barDemo'*/
                ,templet: function(d){
                var url = "/privileges/index/" + d.MenuId
                var viewBtn='<a href='+ url +' class="layui-btn layui-btn-primary layui-btn-xs" lay-event="add">查看</a>';
                var delBtn='<a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>';
                return viewBtn+delBtn;
            }
            }
        ]]
        ,page:false
    });
    treeGrid.on('tool('+tableId+')',function (obj) {
        if(obj.event === 'del'){//删除行
            del(obj);
        }else if(obj.event==="add"){//添加行
            add(obj.data);
        }
    });
});