layui.define(function(exports){ 
//提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
  var obj = {
    	setToken : function(token){
			var local = window.localStorage;
			// 储存了token值
			local.login = token;
			if(! local.token) return false;
			return true;
		}, 
		getToken : function(){
			var local = window.localStorage;
			// 判断是否储存了token值
			var token = local.login;
			if(! token) return false;
			return token;
		},
		say : function(){
			alert('hello...')
		}
  };
  
  //输出test接口
  exports('tools', obj);
});    