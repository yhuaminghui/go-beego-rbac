layui.define(function(exports){ 
//提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
	var ajax = {
		base : function(response){ // 基础处理
			return ajaxInit(response);
		},
		skip : function(response,url){ // 跳转
			var res = ajaxInit(response);

			if(res && url != undefined)
			{
				setTimeout(function(){
					window.location.href = url;
				},1000)
			}
		},
		flush : function(response,isParent){ // 刷新 isparent 为true时刷新父级页面 否则刷新父级页面
			var res = ajaxInit(response);
			if(res){
				// 1s后刷新
				setTimeout(function(){
					if(isParent == true){
						window.parent.location.reload();
					}
					window.location.reload();
				},1000)
			}
		},
		close : function(response,isAll){ // 关闭 isAll 为true时关闭所有 否则关闭当前层
			var res = ajaxInit(response);
			if(res){
				// 1s后关闭
				setTimeout(function(){
					if(isAll == true){
						layer.closeAll();
					}
					var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
		            parent.layer.close(index); //再执行关闭
				},1000)
			}
		},
		
	};
  
  //输出test接口
  exports('ajaxResponse', ajax);
});  

// 基础ajax 提示信息 暂无效果
function ajaxInit(response){
	var error = response.error;
	var msg = response.msg;
	layer.msg(msg);
	if(error != 0){
		return false;
	}
	return true;
}  