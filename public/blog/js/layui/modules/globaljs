/**
 * global
 */

layui.define(['code', 'element', 'table', 'util', 'carousel', 'laytpl'], function(exports){
  var $ = layui.jquery
  ,element = layui.element
  ,layer = layui.layer
  ,form = layui.form
  ,util = layui.util
  ,carousel = layui.carousel
  ,laytpl = layui.laytpl
  ,device = layui.device()
  ,url = layui.url()

  ,$win = $(window), $body = $('body');

  //ban iframe
  ;!function(){self!==parent&&(location.href="//www.baidu.com/")}();

  //阻止 IE7 以下访问
  if(device.ie && device.ie < 8){
    layer.alert('Layui 最低支持 IE8，而您当前使用的是古老的 IE'+ device.ie + '，体验将会不佳！');
  }

  var home = {
    //获取高级浏览器
    getBrowser: function(){
      var ua = navigator.userAgent.toLocaleLowerCase()
      ,mimeType = function(option, value) {
        var mimeTypes = navigator.mimeTypes;
        for (var key in mimeTypes) {
          if (mimeTypes[key][option] && mimeTypes[key][option].indexOf(value) !== -1) {
            return true;
          }
        }
        return;
      };
      if(ua.match(/chrome/)){
        if(mimeType('type', '360') || mimeType('type', 'sogou')) return;
        if(ua.match(/edg\//)) return 'edge';
        return 'chrome'
      } else if(ua.match(/firefox/)){
        return 'firefox';
      }
      
      return;
    }
    //打开新窗口
    ,openWin: function(options){
      var win;
      options = options || {};
      win = options.window || window.open((options.url || ''), options.target, options.specs);
      if(options.url) return;
      win.document.open('text/html', 'replace');
      win.document.write(options.content || '');
      win.document.close();
    }
    ,devHost: 'dev.layuion.com' //+ location.hostname.match(/\w+\.com$/)
  };

  var elemHome = $('#LAY_home');
  var local = layui.data('layui');
  var browser = home.getBrowser();

  //初始弹窗
  layer.ready(function(){
    //升级提示
    if(local.version){
      var vLocal = local.version.split('.').reverse()
      ,vNow = layui.v.split('.').reverse()
      ,updatable;

      //遍历版本号是否存在升级
      layui.each(vLocal, function(i, v){
        if(parseInt(vNow[i]) > parseInt(v)){
          updatable = true;
        } else if(parseInt(vNow[i]) < parseInt(v)) {
          updatable = false;
        }
      });
      
      if(!updatable) return;

      layer.open({
        type: 1
        ,title: '更新提示' //不显示标题栏
        ,closeBtn: false
        ,area: '300px;'
        ,shade: false
        ,offset: '8px'
        ,id: 'LAY_updateNotice' //设定一个id，防止重复弹出
        ,btn: ['更新日志', '我知道了']
        ,btnAlign: 'c'
        ,moveType: 1 //拖拽模式，0或者1
        ,content: ['<div class="layui-text">'
          ,'Layui 已发布新版本：：<strong style="padding-right: 10px; color: #fff;">v'+ layui.v + '</strong>'
        ,'</div>'].join('')
        ,skin: 'layui-layer-notice'
        ,yes: function(index){
          layer.close(index);
          setTimeout(function(){
            location.href = ((
              url.pathname[0] === 'v2' 
              || location.hostname === 'layui.github.io'
            ) ? '/v2' : '') 
            +'/docs/base/changelog.html';
          }, 500);
        }
        ,end: function(){
          layui.data('layui', {
            key: 'version'
            ,value: layui.v
          });
        }
      });
    }
    layui.data('layui', {
      key: 'version'
      ,value: layui.v
    });
  });

  //头部搜索
  ;!function(){
    var elemComponentSelect = $(['<select lay-search lay-filter="component">'
      ,'<option value="">搜索组件模块</option>'
      ,'<option value="element/layout.html">grid 栅格布局</option>'
      ,'<option value="element/layout.html#admin">layout admin 布局</option>'
      ,'<option value="element/color.html">color 颜色</option>'
      ,'<option value="element/icon.html">iconfont 字体图标</option>'
      ,'<option value="base/element.html#css">font 字体大小 颜色</option>'
      ,'<option value="element/anim.html">animation 动画</option>'
      ,'<option value="element/button.html">button 按钮</option>'
      ,'<option value="element/form.html">form 表单组</option>'
      ,'<option value="element/form.html#input">input 输入框</option>'
      ,'<option value="element/form.html#select">select 下拉选择框</option>'
      ,'<option value="element/form.html#checkbox">checkbox 复选框</option>'
      ,'<option value="element/form.html#switch">switch 开关</option>'
      ,'<option value="element/form.html#radio">radio 单选框</option>'
      ,'<option value="element/form.html#textarea">textarea 文本域</option>'
      ,'<option value="element/nav.html">nav 导航菜单</option>'
      ,'<option value="element/menu.html">menu 基础通用菜单</option>'
      ,'<option value="element/nav.html#breadcrumb">breadcrumb 面包屑</option>'
      ,'<option value="element/tab.html">tabs 选项卡</option>'
      ,'<option value="element/progress.html">progress 进度条</option>'
      ,'<option value="element/panel.html">panel 面板 / card</option>'
      ,'<option value="element/collapse.html">collapse 折叠面板/手风琴</option>'
      ,'<option value="element/table.html">table 表格元素</option>'
      ,'<option value="element/badge.html">badge 徽章</option>'
      ,'<option value="element/timeline.html">timeline 时间线</option>'
      ,'<option value="element/auxiliar.html#blockquote">blockquote 引用块</option>'
      ,'<option value="element/auxiliar.html#fieldset">fieldset 字段集</option>'
      ,'<option value="element/auxiliar.html#hr">hr 分割线</option>'
      
      ,'<option value="modules/layer.html">layer 弹出层/弹窗综合</option>'
      ,'<option value="modules/laydate.html">laydate 日期时间选择器</option>'
      ,'<option value="modules/laypage.html">laypage 分页</option>'
      ,'<option value="modules/laytpl.html">laytpl 模板引擎</option>'
      ,'<option value="modules/table.html">table 数据表格</option>'
      ,'<option value="modules/form.html">form 表单模块</option>'
      ,'<option value="modules/upload.html">upload 文件/图片上传</option>'
      ,'<option value="modules/dropdown.html">dropdown 下拉菜单</option>'
      ,'<option value="modules/dropdown.html#contextmenu">contextmenu 右键菜单</option>'
      ,'<option value="modules/transfer.html">transfer 穿梭框</option>'
      ,'<option value="modules/tree.html">tree 树形菜单</option>'
      ,'<option value="modules/colorpicker.html">colorpicker 颜色选择器</option>'
      ,'<option value="modules/element.html">element 常用元素操作</option>'
      ,'<option value="modules/slider.html">slider 滑块</option>'
      ,'<option value="modules/rate.html">rate 评分</option>'
      ,'<option value="modules/carousel.html">carousel 轮播/跑马灯</option>'
      ,'<option value="modules/layedit.html">layedit 富文本编辑器</option>'
      ,'<option value="modules/flow.html">flow 信息流/图片懒加载</option>'
      ,'<option value="modules/util.html">util 工具集</option>'
      ,'<option value="modules/code.html">code 代码文本行修饰</option>'

      ,'<option value="https://'+ home.devHost +'/themes/layim/">layim</option>'
      ,'<option value="https://'+ home.devHost +'/themes/layuiAdmin/">layuiadmin</option>'
    ,'</select>'
    ,'<i class="layui-icon layui-icon-search"></i>'].join(''));

    $('.component').append(elemComponentSelect);
    form.render('select', 'LAY-site-header-component');

    //搜索组件
    form.on('select(component)', function(data){
      var value = data.value;
      if(/^(\/|http)/.test(value)){
        home.openWin({
          url: value
          ,target: '_blank'
        });
        data.elem.value = '';
      } else {
        location.href = (url.pathname[0] === 'v2' ? '/v2' : '') + ('/docs/'+ value);
      }
      
    });
  }();

  // tips notice
  var tipsNotice = function(options, elemParter){
    var local = layui.data('layui');

    options = options || {};

    if(device.mobile) return;

    //是否不显示 tips
    var keyName = 'notice_topnav_'+ options.key
    ,notParter = local[keyName] && new Date().getTime() - local[keyName] < (options.tipsInterval || 1000*60*30); //默认 30 分钟出现一次

    if(!options.tips) layer.close(layer.tipsIndex);

    if(!notParter && options.tips){
      var tipsIndex = layer.tipsIndex = layer.tips(
        ['<a href="'+ options.url +'" target="_blank" style="display: block; line-height: 30px; padding: 10px; text-align: center; font-size: 14px; background-image: linear-gradient(to right,#8510FF,#D025C2,#FF8B2D,#FF0036); color: #fff; '+ (options.tipsCss || '') +'">' 
          ,options.desc || ''
        ,'</a>'].join('')
        ,elemParter
        ,{
          tips: (options.tipsStyle ? new Function('return '+ options.tipsStyle)() : [3, '#9F17E9'])
          ,skin: 'layui-hide-xs'
          ,maxWidth: 320
          ,time: 0
          ,anim: 5
          ,tipeMore: true
          ,success: function(layero, index){
            layero.find('.layui-layer-content').css({
              'padding': 0
            });
            layero.find('a').on('click', function(){
              elemParter.trigger('click');
            });

            //隐藏小箭头
            var tipsG = layero.find('.layui-layer-TipsG');
            if(tipsG.css('left') !== '5px'){
              tipsG.hide();
            }

            //移动端样式
            if(elemParter.parent().css('display') === 'none'){
              layero.css({
                left: '50%'
                ,top: '80px'
                ,'margin-left': -(layero.width()/2)
              });
              tipsG.hide();
            }
          }
        }
      )
      //点击链接
      elemParter.on('click', function(){
        layui.data('layui', {
          key: keyName
          ,value: new Date().getTime()
        });
        layer.close(tipsIndex);
      });
    }

  };

  // spread
  var spread = {
    // 头部 notice
    headerNotice: function(data){
      var noticeElem = $('.site-notice');

      if(device.mobile || !noticeElem[0]) return;
      if(!browser) return;

      data = data || [];
      data = layui.sort(data, 'sort', true); //优先级排序

      var tpl = ['{{# if(d.length > 0){ }}'
      ,'<div class="layui-carousel" id="layui-spm-event-parter" lay-filter="site-top-carousel">'
        ,'<div carousel-item>'
          ,'{{# layui.each(d, function(index, v){ ' 
            ,'var tg = v.ad ? "tg" : "";'
            ,'var style = v.tipsCss || "";'
          ,'}}'
          ,'<div>'
            ,'<a href="{{ v.url }}" target="_blank" class="{{ tg }} tg-{{ v.key }}" data-tips="{{ v.tips }}">'
              ,'<cite style="{{ style }}">{{ v.title }} {{# if(v.hot){ }}<span class="layui-badge-dot" style="margin-top: -5px;"></span>{{# } }}</cite>'
            ,'</a>'
            ,'<style>'
              ,'{{# if(v.ad){ }} .site-notice a.tg-{{ v.key }} cite{padding-right:25px;} {{# } }}'
              ,'{{# if(v.ad){ }}.site-notice a.tg-{{ v.key }}:after{content:"{{ v.ad }}"} {{# } }}'
            ,'</style>'
          ,'</div>'
          ,'{{# }); }}'
        ,'</div>'
      ,'</div>'
      ,'{{# } }}'].join('');

      laytpl(tpl).render(data, function(html){
        var elem = '.site-notice .layui-carousel';
        noticeElem.html(html);

        //轮播实例
        carousel.render({
          elem: elem
          ,width: '100%' //设置容器宽度
          ,height: '100%'
          ,arrow: 'none' //始终显示箭头
          ,indicator: 'none' //指示器位置
          ,anim: 'fade' //切换动画方式
          ,interval: 5000 //自动切换的时间间隔
        });

        tipsNotice(data[0], $(elem).children('div').children('div').eq(0).find('a'));
        carousel.on('change(site-top-carousel)', function(obj){
          tipsNotice(data[obj.index], obj.item.find('a'));
        });
        
      });
    }

    // 头部动态导航
    ,headerNav: function(data){
      var elemNavTop = $('#LAY_NAV_TOP')

      if(!(browser === 'chrome' || browser === 'firefox')) return;
      if(!elemNavTop[0]) return;

      data = data || [];
      data = data[0] || {};
      
      var content = data.content;
      if(!content) return;

      elemNavTop.append(content);

      elemNavTop.find('.layui-nav-bar').remove();
      elemNavTop.find('.layui-nav-item').off('mouseenter').off('mouseleave')
      element.render('nav');
    }

    // 弹出层通知
    ,popupNotive: function(data){
      data = data || [];
      data = data[0] || {};
      
      var content = data.content;
      if(!content) return;

      var hasClickNotice = local.popup_notice && new Date().getTime() - local.popup_notice < (data.tipsInterval || 1000*60*60*24*3);

      if(hasClickNotice) return;

      setTimeout(function(){
        layer.open({
          type: 1
          ,title: data.title || '公告'
          ,area: device.mobile ? ['90%', '90%'] : ['800px', '520px']
          ,shade: false
          //,offset: 'b'
          ,id: 'LAY_Notice' //设定一个id，防止重复弹出
          ,skin: 'site-popup-notice'
          ,resize: false
          ,content: content
          ,success: function(layero, index){
            layero.find('a').on('click', function(){
              layer.close(index);
            });
          }
          ,end: function(){
            layui.data('layui', {
              key: 'popup_notice'
              ,value: new Date().getTime()
            });
          }
        });
      }, 500);

    }
  }

  if(!0){
    $.get('https://'+ home.devHost +'/get/cors/spread/', function(res){
      for(var key in res){
        spread[key] && spread[key](res[key]);
      }
    }, 'jsonp');
  }


  //设置菜单状态
  (function(){
    var docsMenus = $('.site-menu .layui-menu-body-title')
    ,demoMenus = $('.site-demo-nav dd')
    ,pathname = url.pathname.join('/');

    // header
    $('.layui-header .layui-nav-item').each(function(){
      var othis = $(this)
      ,pathname = url.pathname.concat();
      if(pathname[0] === 'v2') pathname.splice(0, 1);
      if(pathname[0] && pathname[0] === othis.data('dir')){
        othis.addClass('layui-this');
        return false;
      }
    });
    
    // docs
    docsMenus.each(function(){
      var othis = $(this)
      ,alink = othis.children('a');
      if(alink.attr('href') === '/'+ pathname){
        othis.parent().addClass('layui-menu-item-checked2');
        return false;
      }
    });

    // demo
    demoMenus.each(function(){
      var othis = $(this)
      ,alink = othis.find('>a');
      if(alink.attr('href') === '/'+ pathname){
        othis.addClass('layui-this');
        return false;
      }
    });

    // demo table
    $('.site-demo-table-nav .layui-nav-item').each(function(){
      var othis = $(this)
      ,alink = othis.find('>a');
      if(alink.attr('href') === '/'+ pathname){
        othis.addClass('layui-this');
        return false;
      }
    });
  })();

  //底部动态信息
  (function(){
    var list = [
      '<a href="https://beian.miit.gov.cn/" target="_blank" rel="nofollow">赣ICP备2021008982号</a>'
      ,'本站托管于<a href="https://gitee.com/" target="_blank">Gitee Pages</a>'
    ]
    ,footerInfo = $('#LAY-footer-info');

    if(/layuion\.com/.test(location.hostname)){
      footerInfo.append(list[0]);
    } else {
      footerInfo.append(list[1]);
    }
  })();


  //展示当前版本
  $('.site-showv').html(layui.v);

  //获取 Github 数据
  var getStars = $('#getStars');
  if(getStars[0]){
    $.get('https://api.github.com/repos/layui/layui', function(res){
      getStars.html(res.stargazers_count);
    }, 'json');
  }
  
  //固定Bar
  util.fixbar({
    showHeight: 60
    ,css: function(){
      if(global.pageType === 'demo'){
        return {bottom: 75}
      }
    }()
  });
  
  //窗口scroll
  (function(){
    var main = $('.site-menu'), scroll = function(){
      var stop = $(window).scrollTop();

      if($(window).width() <= 992) return;
      var bottom = $('.footer').offset().top - $(window).height();

      if(stop > 60){ //211
        if(!main.hasClass('site-fix')){
          main.addClass('site-fix').css({
            width: main.parent().width()
          });
        }
      }else {     
        if(main.hasClass('site-fix')){
          main.removeClass('site-fix').css({
            width: 'auto'
          });
        }
      }
      stop = null;
    };
    scroll();
    $(window).on('scroll', scroll);
  })();

  //示例页面滚动
  $(window).on('scroll', function(){
    var elemTips = $('.layui-table-tips');
    if(elemTips[0]) elemTips.remove();

    if($('.layui-layer')[0]){
      layer.closeAll('tips');
    }
  });
  
  //代码修饰
  layui.code({
    elem: 'pre'
  });
  
  //目录
  var siteDir = $('.site-dir');
  if(siteDir[0] && $(window).width() > 750){
    layer.ready(function(){
      layer.open({
        type: 1
        ,content: siteDir
        ,skin: 'layui-layer-dir'
        ,area: 'auto'
        ,maxHeight: $(window).height() - 300
        ,title: '目录'
        ,closeBtn: false
        ,offset: 'r'
        ,shade: false
        ,success: function(layero, index){
          layer.style(index, {
            marginLeft: -15
          });
        }
      });
    });
    siteDir.find('li').on('click', function(){
      var othis = $(this);
      othis.find('a').addClass('layui-this');
      othis.siblings().find('a').removeClass('layui-this');
    });
  }

  //在 textarea 焦点处插入字符
  var focusInsert = function(str){
    var start = this.selectionStart
    ,end = this.selectionEnd
    ,offset = start + str.length

    this.value = this.value.substring(0, start) + str + this.value.substring(end);
    this.setSelectionRange(offset, offset);
  };

  //演示页面
  $('body').on('keydown', '#LAY_editor, .site-demo-text', function(e){
    var key = e.keyCode;
    if(key === 9 && window.getSelection){
      e.preventDefault();
      focusInsert.call(this, '  ');
    }
  });

  var editor = $('#LAY_editor')
  ,iframeElem = $('#LAY_demo')
  ,demoCodes = $('#LAY_demoCodes')[0]
  ,runCodes = function(){
    if(!iframeElem[0]) return;
    var html = editor.val();
    iframeElem[0].srcdoc = html
  };
  $('#LAY_demo_run').on('click', runCodes), runCodes();


  //让选中的菜单保持在可视范围内
  util.toVisibleArea({
    scrollElem: $('.layui-side-scroll').eq(0)
    ,thisElem: $('.site-demo-nav').find('dd.layui-this')
  });

  util.toVisibleArea({
    scrollElem: $('.layui-side-scroll').eq(1)
    ,thisElem: $('.site-demo-table-nav').find('li.layui-this')
  });

  //查看代码
  $(function(){
    var demoCodeItem = $('.site-demo-body>.layui-tab-item').eq(1)
    ,demoCode = $('<textarea class="layui-border-box site-demo-text site-demo-code" id="LAY_democode" spellcheck="false" readonly></textarea>');
    
    demoCode.val([
      '<!DOCTYPE html>'
      ,'<html>'
      ,'<head>'
        ,'  <meta charset="utf-8">'
        ,'  <title>示例演示</title>'
        ,'  <meta name="renderer" content="webkit">'
        ,'  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">'
        ,'  <meta name="viewport" content="width=device-width, initial-scale=1">'
        ,'  <!-- 注意：项目正式环境请勿引用该地址 -->'
        ,'  <link href="//unpkg.com/layui@'+ layui.v +'/dist/css/layui.css" rel="stylesheet">'
      ,'</head>'
      ,'<body>'
      ,global.preview
      ,'<!-- 注意：项目正式环境请勿引用该地址 -->'
      ,'<script src="//unpkg.com/layui@'+ layui.v +'/dist/layui.js"></script>'
      ,$('#LAY_democodejs').html()
      ,'</body>\n</html>'
    ].join('\n'));

    demoCodeItem.html(demoCode);
  });

  //点击查看代码选项
  element.on('tab(demoTitle)', function(obj){
    if(obj.index === 1){
      if(device.ie && device.ie < 9){
        layer.alert('强烈不推荐你通过ie8/9 查看代码！因为，所有的标签都会被格式成大写，且没有换行符，影响阅读');
      }
    }
  })


  //手机设备的简单适配
  var treeMobile = $('.site-tree-mobile')
  ,shadeMobile = $('.site-mobile-shade')

  treeMobile.on('click', function(){
    $('body').addClass('site-mobile');
  });

  shadeMobile.on('click', function(){
    $('body').removeClass('site-mobile');
  });

  //获取文档图标数据
  home.getIconData = function(){
    var $ = layui.$
    ,iconData = []
    ,iconListElem = $('.site-doc-icon li');

    iconListElem.each(function(){
      var othis = $(this);
      iconData.push({
        title: $.trim(othis.find('.doc-icon-name').text())
        ,fontclass: $.trim(othis.find('.doc-icon-fontclass').text())
        ,unicode: $.trim(othis.find('.doc-icon-code').html())
      });
    });
    
    $('.site-h1').html('<textarea style="width: 100%; height: 600px;">'+ JSON.stringify(iconData) + '</textarea>');
  };

  exports('global', home);
});