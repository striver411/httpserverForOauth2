<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta charset="utf-8" />
		<title>Dashboard - DeepShare Admin</title>

		<meta name="description" content="overview &amp; stats" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />

		<!-- bootstrap & fontawesome -->
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/bootstrap.css" />
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/font-awesome.css" />

		<!-- page specific plugin styles -->

		<!-- text fonts -->
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/ace-fonts.css" />
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/select2.css" />
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/daterangepicker.css" />
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/datepicker.css" />

		<!-- ace styles -->
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/ace.css" class="ace-main-stylesheet" id="main-ace-style" />

		<!--[if lte IE 9]>
			<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/ace-part2.css" class="ace-main-stylesheet" />
		<![endif]-->

		<!--[if lte IE 9]>
		  <link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/ace-ie.css" />
		<![endif]-->

		<!-- inline styles related to this page -->

		<!-- ace settings handler -->
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace-extra.js"></script>

		<!-- HTML5shiv and Respond.js for IE8 to support HTML5 elements and media queries -->

		<!--[if lte IE 8]>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/html5shiv.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/respond.js"></script>
		<![endif]-->
		<style>
			.dropdown-content li.active{
				border-left: 4px solid #00adef;
			}
		</style>
	</head>

	<body class="no-skin">
		<!-- #section:basics/navbar.layout -->
		<div id="navbar" class="navbar navbar-default" style="background-color: #404040">
			<script type="text/javascript">
				try{ace.settings.check('navbar' , 'fixed')}catch(e){}
			</script>

			<div class="navbar-container" id="navbar-container">
				<!-- #section:basics/sidebar.mobile.toggle -->
				<button type="button" class="navbar-toggle menu-toggler pull-left" id="menu-toggler" data-target="#sidebar">
					<span class="sr-only">Toggle sidebar</span>

					<span class="icon-bar"></span>

					<span class="icon-bar"></span>

					<span class="icon-bar"></span>
				</button>

				<!-- /section:basics/sidebar.mobile.toggle -->
				<div class="navbar-header pull-left">
					<!-- #section:basics/navbar.layout.brand -->
					<a href="#" class="navbar-brand">
						<small>
							<i class="fa fa-leaf"></i>
							DeepShare Admin
						</small>
					</a>

					<!-- /section:basics/navbar.layout.brand -->

					<!-- #section:basics/navbar.toggle -->

					<!-- /section:basics/navbar.toggle -->
				</div>

				<!-- #section:basics/navbar.dropdown -->
				<div class="navbar-buttons navbar-header pull-right" role="navigation">
					<ul class="nav ace-nav">
						<li class="green">
							<a data-toggle="dropdown" class="dropdown-toggle" href="#">
								<span><b id="app-title"></b></span>
								<i class="ace-icon fa fa-chevron-down icon-animated-vertical"></i>
							</a>

							<ul class="dropdown-menu-right dropdown-navbar dropdown-menu dropdown-caret dropdown-close">
								<li class="dropdown-header">
									<i class="ace-icon fa fa-check"></i>
									3 Apps&nbsp;&nbsp;&nbsp;&nbsp;
									<i class="ace-icon fa fa-times"></i>
									1 App
								</li>

								<li class="dropdown-content">
									<ul class="dropdown-menu dropdown-navbar" id="applist">
									</ul>
								</li>

								<li class="dropdown-footer">
									<a href="/addapp">
										添加应用
										<i class="ace-icon fa fa-arrow-right"></i>
									</a>
								</li>
							</ul>
						</li>
						<li class="grey">
							<a data-toggle="dropdown" href="#" class="dropdown-toggle">
								<img class="nav-user-photo" src="<?php echo Yii::app()->baseUrl; ?>/assets/avatars/avatar2.png" alt="<?php echo Yii::app()->user->id; ?>'s Photo" />
								<span class="user-info">
									<small>欢迎,</small>
									<?php echo Yii::app()->user->name; ?>
								</span>

								<i class="ace-icon fa fa-caret-down"></i>
							</a>

							<ul class="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-close">
								<li>
									<a href="#">
										<i class="ace-icon fa fa-cog"></i>
										Settings
									</a>
								</li>

								<li>
									<a href="/profile">
										<i class="ace-icon fa fa-user"></i>
										Profile
									</a>
								</li>

								<li class="divider"></li>

								<li>
									<a href="<?php echo Yii::app()->baseUrl; ?>/logout">
										<i class="ace-icon fa fa-power-off"></i>
										Logout
									</a>
								</li>
							</ul>
						</li>

						<!-- /section:basics/navbar.user_menu -->
					</ul>
				</div>

				<!-- /section:basics/navbar.dropdown -->
			</div><!-- /.navbar-container -->
		</div>

		<!-- /section:basics/navbar.layout -->
		<div class="main-container" id="main-container">
			<script type="text/javascript">
				try{ace.settings.check('main-container' , 'fixed')}catch(e){}
			</script>

			<!-- #section:basics/sidebar -->
			<div id="sidebar" class="sidebar                  responsive">
				<script type="text/javascript">
					try{ace.settings.check('sidebar' , 'fixed')}catch(e){}
				</script>


				<ul class="nav nav-list">
					<li class="active highlight">
						<a href="/index">
							<i class="menu-icon fa fa-tachometer"></i>
							<span class="menu-text"> Dashboard </span>
						</a>
						<b class="arrow"></b>
					</li>

					<li class="active">
						<a href="/index">
							<i class="menu-icon fa fa-list-alt"></i>
							<span class="menu-text"> 数据概览 </span>
						</a>
						<b class="arrow"></b>
					</li>
					<li class="">
						<a href="/addapp">
							<i class="menu-icon fa fa-tag"></i>
							<span class="menu-text"> 添加应用 </span>
						</a>
						<b class="arrow"></b>
					</li>
					<li class="">
						<a href="/profile">
							<i class="menu-icon fa fa-desktop"></i>
							<span class="menu-text"> 账号管理 </span>
						</a>
						<b class="arrow"></b>
					</li>
					<li class="">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-share-alt"></i>
							<span class="menu-text"> Referrals </span>
	
							<b class="arrow fa fa-angle-down"></b>
						</a>
	
						<b class="arrow"></b>
	
						<ul class="submenu">
							<li class="" id="li-overview">
								<a href="/referrals?overview">
									<i class="menu-icon fa fa-caret-right"></i>
									Overview
								</a>
	
								<b class="arrow"></b>
							</li>
							<li class="" id="li-rules">
								<a href="/referrals?rules">
									<i class="menu-icon fa fa-caret-right"></i>
									Reward Rules
								</a>
	
								<b class="arrow"></b>
							</li>
							<li class="" id="li-influencers">
								<a href="/referrals?influencers">
									<i class="menu-icon fa fa-caret-right"></i>
									Influencers
								</a>
	
								<b class="arrow"></b>
							</li>
						</ul>
					</li>
					
					<li class="">
						<a href="/marketing">
							<i class="menu-icon fa fa-link "></i>
							<span class="menu-text"> Marketing </span>
						</a>
						<b class="arrow"></b>
					</li>
				</ul><!-- /.nav-list -->

				<!-- #section:basics/sidebar.layout.minimize -->
				<div class="sidebar-toggle sidebar-collapse" id="sidebar-collapse">
					<i class="ace-icon fa fa-angle-double-left" data-icon1="ace-icon fa fa-angle-double-left" data-icon2="ace-icon fa fa-angle-double-right"></i>
				</div>

				<!-- /section:basics/sidebar.layout.minimize -->
				<script type="text/javascript">
					try{ace.settings.check('sidebar' , 'collapsed')}catch(e){}
				</script>
			</div>

			<!-- /section:basics/sidebar -->
			<div class="main-content">
				<div class="main-content-inner">
					<!-- #section:basics/content.breadcrumbs -->
					<div class="breadcrumbs" id="breadcrumbs">
						<script type="text/javascript">
							try{ace.settings.check('breadcrumbs' , 'fixed')}catch(e){}
						</script>

						<ul class="breadcrumb">
							<li>
								<i class="ace-icon fa fa-home home-icon"></i>
								<a href="#">Home</a>
							</li>
							<li class="active">Dashboard</li>
						</ul><!-- /.breadcrumb -->

						<!-- #section:basics/content.searchbox -->
						<!-- <div class="nav-search" id="nav-search">
							<form class="form-search">
								<span class="input-icon">
									<input type="text" placeholder="Search ..." class="nav-search-input" id="nav-search-input" autocomplete="off" />
									<i class="ace-icon fa fa-search nav-search-icon"></i>
								</span>
							</form>
						</div> --><!-- /.nav-search -->
						<!-- /section:basics/content.searchbox -->
					</div>

					<!-- /section:basics/content.breadcrumbs -->
					<div class="page-content">

						<!-- /section:settings.box -->
						<div class="page-header">
							<h1>
								Dashboard
								<small>
									<i class="ace-icon fa fa-angle-double-right"></i>
									数据概览
								</small>
							</h1>
						</div><!-- /.page-header -->

						<div class="row">
							<div class="col-xs-12">
								
								<div class="row">
									<div class="space-6"></div>
									<div class="col-sm-12 infobox-container">
										<!-- #section:pages/dashboard.infobox -->
										<div class="infobox infobox-green">
											<!-- #section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-chart">
												<span class="sparkline" data-values="196,128,202,177,154,94,100,170,224"></span>
											</div>

											<!-- /section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-data">
												<span class="infobox-data-number" id="install_today"></span>
												<div class="infobox-content"><b>今日安装</b></div>
											</div>

											<div class="badge badge-success">
												7.2%
												<i class="ace-icon fa fa-arrow-up"></i>
											</div>
										</div>
										<div class="infobox infobox-green">
											<!-- #section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-chart">
												<span class="sparkline" data-values="154,94,100,170,224,196,128,202,177"></span>
											</div>
<!-- 											<div class="easy-pie-chart percentage" data-percent="42" data-size="46"> -->
<!-- 												<span class="percent">42</span>% -->
<!-- 											</div> -->

											<!-- /section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-data">
												<span class="infobox-data-number" id="open_today"></span>
												<div class="infobox-content"><b>今日使用</b></div>
											</div>

											<div class="badge badge-success">
												10%
												<i class="ace-icon fa fa-arrow-up"></i>
											</div>
										</div>
<!-- 										<div class="space-6"></div> -->
										<div class="infobox infobox-blue">
											<!-- #section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-chart">
												<span class="sparkline" data-values="196,128,202,177,154,94,100,170,224"></span>
											</div>

											<!-- /section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-data">
												<span class="infobox-data-number" id="install_last_day"></span>
												<div class="infobox-content"><b>昨日安装</b></div>
											</div>

											<div class="badge badge-success">
												7.2%
												<i class="ace-icon fa fa-arrow-up"></i>
											</div>
										</div>
										<div class="infobox infobox-blue">
											<!-- #section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-chart">
												<span class="sparkline" data-values="154,94,100,170,224,196,128,202,177"></span>
											</div>
<!-- 											<div class="easy-pie-chart percentage" data-percent="42" data-size="46"> -->
<!-- 												<span class="percent">42</span>% -->
<!-- 											</div> -->

											<!-- /section:pages/dashboard.infobox.sparkline -->
											<div class="infobox-data">
												<span class="infobox-data-number" id="open_last_day"></span>
												<div class="infobox-content"><b>昨日使用</b></div>
											</div>

											<div class="badge badge-success">
												10%
												<i class="ace-icon fa fa-arrow-up"></i>
											</div>
										</div>
									</div>

									<div class="vspace-12-sm"></div>
								</div><!-- /.row -->

								<div class="hr hr32 hr-dotted"></div>
								<div class="row">
									<div class="col-sm-6">
										<ul class="pagination">
											<li>
												<a href="#">按时统计</a>
											</li>
											<li class="active">
												<a href="#">按天统计</a>
											</li>
											<li >
												<a href="#">按周统计</a>
											</li>
											<li>
												<a href="#">按月统计</a>
											</li>
										</ul>
									</div>
									<div class="col-sm-3" style="margin-top: 24px;text-align: right;
								font-size: 14px;font-weight: bold;color: #2283c5">
									选择时段:</div>
									<div class="col-sm-3">
										<div class="input-daterange input-group" style="margin-top: 20px">
											<input type="text" class="input-sm form-control" name="start"/>
											<span class="input-group-addon">
												<i class="fa fa-exchange"></i>
											</span>

											<input type="text" class="input-sm form-control" name="end"/>
										</div>
									</div>
								</div>
								<div class="row">
									<div class="col-sm-12">
										<div id="main" style="height:400px;"></div>
									</div>
								</div>
								<!-- PAGE CONTENT ENDS -->
							</div><!-- /.col -->
						</div><!-- /.row -->
					</div><!-- /.page-content -->
				</div>
			</div><!-- /.main-content -->

			<div class="footer">
				<div class="footer-inner">
					<!-- #section:basics/footer -->
					<div class="footer-content">
						<span class="bigger-120">
							<span class="blue bolder">DeepShare</span>
							Application &copy; 2015
						</span>

					</div>

					<!-- /section:basics/footer -->
				</div>
			</div>

			<a href="#" id="btn-scroll-up" class="btn-scroll-up btn btn-sm btn-inverse">
				<i class="ace-icon fa fa-angle-double-up icon-only bigger-110"></i>
			</a>
		</div><!-- /.main-container -->

		<!-- basic scripts -->

		<!--[if !IE]> -->
		<script type="text/javascript">
			window.jQuery || document.write("<script src='<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery.js'>"+"<"+"/script>");
		</script>

		<!-- <![endif]-->

		<!--[if IE]>
<script type="text/javascript">
 window.jQuery || document.write("<script src='<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery1x.js'>"+"<"+"/script>");
</script>
<![endif]-->
		<script type="text/javascript">
			if('ontouchstart' in document.documentElement) document.write("<script src='<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery.mobile.custom.js'>"+"<"+"/script>");
		</script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/bootstrap.js"></script>

		<!-- page specific plugin scripts -->

		<!--[if lte IE 8]>
		  <script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/excanvas.js"></script>
		<![endif]-->
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery-ui.custom.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery.ui.touch-punch.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery.easypiechart.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/jquery.sparkline.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/flot/jquery.flot.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/flot/jquery.flot.pie.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/flot/jquery.flot.resize.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/select2.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/date-time/bootstrap-datepicker.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/date-time/daterangepicker.js"></script>
		
		<!-- ace scripts -->
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/elements.scroller.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/elements.wizard.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/elements.aside.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.ajax-content.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.sidebar.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.sidebar-scroll-1.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.widget-box.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.widget-on-reload.js"></script>

		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/echarts.js"></script>
		<script type="text/javascript">
	        
	    </script>
		<!-- inline scripts related to this page -->
		<script type="text/javascript">
			function addDate(dadd){
				var a = new Date();
				a = a.valueOf();
				a = a + dadd * 24 * 60 * 60 * 1000;
				a = new Date(a);
				return a;
			}
			function getData(appid) {
				$.ajax({
                    type: "POST",
                    dataType: "json",
                    url: "/applist",
                    data:{    
				    	"data1":"从前台传到后台的值",
				    	"data2":"可以在这里添加",
				    	"appid":appid
				    }, 
                    success: function (result) {
                        $("#app-title").html(result.appname);
                        $("#open_today").html(result.appstats.open_today);
                        $("#install_today").html(result.appstats.install_today);
                        $("#open_last_day").html(result.appstats.open_last_day);
                        $("#install_last_day").html(result.appstats.install_last_day);

                        for(var a in result.applist) {
							$("#applist").append('<li class='+(result.applist[a].appname==result.appname?'active':'')+'>'+
								'<a href="#"><div class="clearfix"><span class="pull-left">'+result.applist[a].appname+'</span>'+
								'<span class="pull-right"><i class="fa fa-check"></i></span></div></a></li>');
                        }
                        var x = new Array(); 
                        var openlist = new Array(); 
                        var installlist = new Array(); 
                        for(var i in result.appstats.data) {
                            x[i] = addDate(i-result.appstats.data.length).getDate();
                        	openlist[i] = result.appstats.data[i].open;
                        	installlist[i] = result.appstats.data[i].install;
                        }

                        require.config({
            	            paths: {
            	            	echarts: '<?php echo Yii::app()->baseUrl; ?>/assets/js'
            	            }
            	        });
            	        require(
            	                [
            	                    'echarts',
            	                    'echarts/chart/line'
            	                ],
            	                function (ec) {
            	                    var myChart = ec.init(document.getElementById('main'));
            	                    var option = {
            	                    	    title : {
            	                    	        text: '统计',
            	                    	        subtext: '示例'
            	                    	    },
            	                    	    tooltip : {
            	                    	        trigger: 'axis'
            	                    	    },
            	                    	    legend: {
            	                    	        data:['新增使用','新增安装']
            	                    	    },
            	                    	    toolbox: {
            	                    	        show : true,
            	                    	        feature : {
            	                    	            dataView : {show: true, readOnly: false},
            	                    	            saveAsImage : {show: true}
            	                    	        }
            	                    	    },
            	                    	    calculable : true,
            	                    	    xAxis : [
            	                    	        {
            	                    	            type : 'category',
            	                    	            boundaryGap : false,
            	                    	            data : x
            	                    	        }
            	                    	    ],
            	                    	    yAxis : [
            	                    	        {
            	                    	            type : 'value',
            	                    	            axisLabel : {
            	                    	                formatter: '{value}'
            	                    	            }
            	                    	        }
            	                    	    ],
            	                    	    series : [
            	                    	        {
            	                    	            name:'新增使用',
            	                    	            type:'line',
            	                    	            data:openlist,
            	                    	            markPoint : {
            	                    	                data : [
            	                    	                    {type : 'max', name: '最大值'},
            	                    	                    {type : 'min', name: '最小值'}
            	                    	                ]
            	                    	            }
            	                    	        },
            	                    	        {
            	                    	            name:'新增安装',
            	                    	            type:'line',
            	                    	            data:installlist,
            	                    	            markPoint : {
            	                    	            	data : [
            		                    	                    {type : 'max', name: '最大值'},
            		                    	                    {type : 'min', name: '最小值'}
            		                    	                ]
            	                    	            }
            	                    	        }
            	                    	    ]
            	                    	};
            	                    	                    
            	                    myChart.setOption(option);
            	                }
            	            );
                    },
                    error: function(data) {
	                    alert("error:"+data.responseText);
                    }
               });
			}
			
			jQuery(function($) {
				getData();
				
				$('.input-daterange').datepicker({autoclose:true,format:'yyyy-mm-dd'});
				
				$(document).delegate('.dropdown-menu a', 'click', function(){
					//TODO：点击后应该触发切换app的action，所有展示数字全部重新获取
					//getData(传入appid);
					var title=$(this).children("div").children("span.pull-left").html();
					$("#app-title").html(title);
					$(this).parent().siblings().removeClass("active");
					$(this).parent().addClass("active");
				}); 
				
				$('[data-rel=tooltip]').tooltip();
				
				$(".select2").css('width','200px').select2({allowClear:true})
				.on('change', function(){
					$(this).closest('form').validate().element($(this));
				}); 
				
				$('.easy-pie-chart.percentage').each(function(){
					var $box = $(this).closest('.infobox');
					var barColor = $(this).data('color') || (!$box.hasClass('infobox-dark') ? $box.css('color') : 'rgba(255,255,255,0.95)');
					var trackColor = barColor == 'rgba(255,255,255,0.95)' ? 'rgba(255,255,255,0.25)' : '#E2E2E2';
					var size = parseInt($(this).data('size')) || 50;
					$(this).easyPieChart({
						barColor: barColor,
						trackColor: trackColor,
						scaleColor: false,
						lineCap: 'butt',
						lineWidth: parseInt(size/10),
						animate: /msie\s*(8|7|6)/.test(navigator.userAgent.toLowerCase()) ? false : 1000,
						size: size
					});
				})
			
				$('.sparkline').each(function(){
					var $box = $(this).closest('.infobox');
					var barColor = !$box.hasClass('infobox-dark') ? $box.css('color') : '#FFF';
					$(this).sparkline('html',
									 {
										tagValuesAttribute:'data-values',
										type: 'bar',
										barColor: barColor ,
										chartRangeMin:$(this).data('min') || 0
									 });
				});
			
			
				$('#recent-box [data-rel="tooltip"]').tooltip({placement: tooltip_placement});
				function tooltip_placement(context, source) {
					var $source = $(source);
					var $parent = $source.closest('.tab-content')
					var off1 = $parent.offset();
					var w1 = $parent.width();
			
					var off2 = $source.offset();
					//var w2 = $source.width();
			
					if( parseInt(off2.left) < parseInt(off1.left) + parseInt(w1 / 2) ) return 'right';
					return 'left';
				}
			
			
				$('.dialogs,.comments').ace_scroll({
					size: 300
			    });
				
			
			})
		</script>

		<!-- the following scripts are used in demo only for onpage help and you don't need them -->
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/assets/css/ace.onpage-help.css" />
		<link rel="stylesheet" href="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/themes/sunburst.css" />

		<script type="text/javascript"> ace.vars['base'] = '..'; </script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/elements.onpage-help.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/assets/js/ace/ace.onpage-help.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/rainbow.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/language/generic.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/language/html.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/language/css.js"></script>
		<script src="<?php echo Yii::app()->baseUrl; ?>/docs/assets/js/language/javascript.js"></script>
	</body>
</html>
