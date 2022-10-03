(function (window) {
    var width = document.getElementById("articleDirectories").offsetWidth
    // 4. 窗口的滚动
    window.onscroll = function () {
       
        // 4.2 判断是否吸顶
        var scrollTop =getScrollTop(); 
        if(scrollTop >=400){
           
			$('articleDirectories').classList.add('fix-articleDirectories')
            document.getElementById("articleDirectories").style.width= width+"px"
       
            
        }else{
			$('articleDirectories').classList.remove('fix-articleDirectories')  
        }
    };

})(window);


// 通用弹窗

	// 标签 分类弹窗
function showDialog(dialogId){
    var $ = layui.$
    layer.open({
        type: 1,
        title: '标签',
        area: ["auto"], //宽高
        resize:false,
        skin: 'layui-layer-rim', 
        closeBtn: 2,
        scrollbar: false,
        shadeClose: true,
        content: $(dialogId).html()
    });
}

function getScrollTop()
{
  var scrollTop=0;
  if(document.documentElement&&document.documentElement.scrollTop)
  {
  scrollTop=document.documentElement.scrollTop;
  }
  else if(document.body)
  {
  scrollTop=document.body.scrollTop;
  }
  return scrollTop;
}

// function init_gotop()
// {
// 	$("#go_top").css("top",$(document).scrollTop()+$(window).height()-80);
// 	$(window).scroll(function(){
// 		$("#go_top").css("top",$(document).scrollTop()+$(window).height()-80);
// 		if($(document).scrollTop()>0)
// 			$("#go_top").fadeIn();
// 		else
// 			$("#go_top").fadeOut();
// 	});	
	
// 	$("#go_top").bind("click",function(){
// 		$("html,body").animate({scrollTop:0},"fast","swing",function(){
// 		});
// 	});
 
// }