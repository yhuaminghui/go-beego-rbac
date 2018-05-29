//config的设置是全局的
layui.config({
  base: '/static/layui_extenstion/' //假设这是你存放拓展模块的根目录
}).extend({ //设定模块别名
  tools: 'tools', //如果 mymod.js 是在根目录，也可以不用设定别名
  ajaxResponse: 'ajax',
});