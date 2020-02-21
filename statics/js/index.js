
var text = $("#f-left");
// text.focus();
var is_review = false;
var is_correct = true;
review_id = 0;
msg_num = 0;

readQus();

// review action
function review_action(query){
	$(".b-body").append("<div class='mWord'><span></span><p>" + query + "</p></div>");
	$(".b-body").append("<div class='typing_loader'></div>");
	$("#content").scrollTop(10000000);	

	var args = {
		type: "get",
		url: "http://localhost:8080/question/review?question=" + query + "&id=review" +review_id.toString(),
		//data: { "appid": "xiaosi", "spoken": text.val() },
		success: function (redata) {
			var my_data = $.parseJSON(redata)

			var array = [my_data.data];
			is_correct = my_data.isCorrect;
			var recommand = my_data.recommand;
			// var recommand = my_data.recommand
			$(".typing_loader").remove();
			//console.log(array);
			for (var i = 0; i < array.length; i++) {
				//console.log(array[i]);
				var result = array[i];
				//var p_id = "p_" + msg_num.toString();
				//$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'> <p id='{0}'>".format(p_id) + result + "</p></div></div>");
				if (review_id == 8){
				$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'>" + result + "<br><button class='ui medium blue button is_button'>无</button></div></div>");
				}
				else if (review_id>=3 && review_id<=7){
				$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'>" + result + "<br><button class='ui medium blue button is_button'>是</button><button class='ui medium blue button is_button'>否</button</div></div>");
				}else{
				$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'>" + result + "</div></div>");
				}	
				if (review_id == 9){
				$(".b-body").append(
					"<div class='rotWord'>\
					<br>\
					<span></span>\
					<div class='rotWord_P'>\
					  您还可以这样问我，或者说“病情自测”：\
					  <br>\
					  <div class='ui blue segment recommand'>\
							{0}\
					  </div>\
					  <div class='ui blue segment recommand'>\
						{1}\
					  </div>\
					  <div class='ui blue segment recommand'>\
						{2}\
					  </div>\
					</div>\
					</div>".format(recommand[0], recommand[1], recommand[2])
				);
					is_review = false;
				}
				//foldText('p#{0}'.format(p_id), 200);
				$("#content").scrollTop(10000000);

				// 若出最终结果则 推荐问题
			}
		}
	}
	ajax(args);
	text.val("");
}

function action() {
	if (text.val() == null || text.val() == "") {
		//text.focus();
		return;
	}
	var query = text.val().trim();
	if (query=="病情自测" || is_review){
		is_review = true;
		if (is_correct){
		review_id = review_id + 1;
		}
		if (review_id > 9){
			review_id = 0;
			is_review = false;
			return;
		}
		review_action(query);
		return;
	}

	$(".b-body").append("<div class='mWord'><span></span><p>" + text.val() + "</p></div>");
	// 加入加载动画
	// msg_num = msg_num + 1;
	$(".b-body").append("<div class='typing_loader'></div>");

	$("#content").scrollTop(10000000);

	var args = {
		type: "get",
		url: "http://localhost:8080/question?question=" + text.val().trim(),
		//data: { "appid": "xiaosi", "spoken": text.val() },
		success: function (redata) {
			var my_data = $.parseJSON(redata)

			var array = [my_data.data];
			var recommand = my_data.recommand

			//if (my_data.data.info.hasOwnProperty("heuristic")) {
			//	for (var i = 0; i < my_data.data.info.heuristic.length; i++) {
			//		array.push(my_data.data.info.heuristic[i]);
			//	}
			//}
			$(".typing_loader").remove();
			//console.log(array);
			for (var i = 0; i < array.length; i++) {
				//console.log(array[i]);
				var result = array[i];
				var p_id = "p_" + msg_num.toString();
				//$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'> <p id='{0}'>".format(p_id) + result + "</p></div></div>");
				$(".b-body").append("<div class='rotWord'><span></span> <div class='rotWord_P'>" + result + "</div></div>");

				// 推荐问题
				$(".b-body").append(
					"<div class='rotWord'>\
					<br>\
					<span></span>\
					<div class='rotWord_P'>\
					  您还可以这样问我，或者说“病情自测”：\
					  <br>\
					  <div class='ui blue segment recommand'>\
							{0}\
					  </div>\
					  <div class='ui blue segment recommand'>\
						{1}\
					  </div>\
					  <div class='ui blue segment recommand'>\
						{2}\
					  </div>\
					</div>\
					</div>".format(recommand[0], recommand[1], recommand[2])
				);

				//foldText('p#{0}'.format(p_id), 200);
				$("#content").scrollTop(10000000);
			}
		}
	}
	ajax(args);
	text.val("");
	// text.focus();

};


String.prototype.format = function () {
	var values = arguments;
	return this.replace(/\{(\d+)\}/g, function (match, index) {
		if (values.length > index) {
			return values[index];
		} else {
			return "";
		}
	});
};

$('#help-icon').popup({
	popup: $('.custom.popup'),
	on: 'click',
	inline: true
});

$("#btn").click(function () {
	action();
	msg_num = msg_num + 1;
	console.log(msg_num);
});
$(document).keydown(function (event) {
	if (event.keyCode == 13) {
		action();
		msg_num = msg_num + 1;
	}
});

// 文字折叠
function foldText(clas, num) {
	var num = num;
	var a = $("<a></a>").on("click", showText).addClass('a-text').text("【展开】");
	var b = $("<a></a>").on("click", showText).addClass('a-text').text("【折叠】");
	var p2_class = 'p2_' + msg_num.toString()
	var p = $("<p></p>").addClass(p2_class);
	var str = $(clas).text();
	$(clas).after(p);
	$('.' + p2_class).hide();
	if (str.length > num) {
		var text = str.substring(0, num) + "...";
		$(clas).html(text).append(a);

	}
	$('.' + p2_class).html(str).append(b);
	function showText() {
		$(this).parent().hide().siblings().show();
	}
}

function ajax(mJson) {
	var type = mJson.type || 'get';
	var url = mJson.url;
	var data = mJson.data;
	var success = mJson.success;
	var error = mJson.error;
	var dataStr = '';

	if (data) {
		var arr = Object.keys(data);
		var len = arr.length;
		var i = 0;

		for (var key in data) {
			dataStr += key + '=' + data[key];

			if (++i < len) {
				dataStr += '&';
			}
		}

		if (type.toLowerCase() == 'get') {
			url += '?' + dataStr;
		}
	}

	//console.log(url);

	var xhr = new XMLHttpRequest();
	xhr.open(type, url, true);
	xhr.setRequestHeader('content-type', 'application/x-www-form-urlencoded');
	xhr.send(null);

	xhr.onreadystatechange = function () {
		if (xhr.readyState == 4) {
			if (xhr.status >= 200 && xhr.status < 300) {
				success && success(xhr.responseText);
			}
			else {
				error && error(xhr.status);
			}
		}
	}
}

window.onload = function () {
	$('#mianze').modal('setting', 'closable', false).modal('show');
}

// 适配手机
// var winHeight = $(window).height();
// $(window).resize(function() {
//     var thisHeight = $(this).height();
//     var keyboardHeight = thisHeight - winHeight;
//     $("#footer").css({ 'bottom': keyboardHeight + 'px' });
// });

text.focus(function () {
	document.querySelector('#footer').scrollIntoView();
})


// $(document).delegate('input, textarea', 'blur', function () {
// 	setTimeout(function () {
// 		jQuery('html').animate({ height: '100.1vh' }, 100, function () {
// 			jQuery(this).animate({ height: '100vh' }, 1)
// 		})
// 	}, 100)
//     $("#content").scrollTop(10000000);
// })
$("input").blur(function () {
	var u = navigator.userAgent;
	var isiOS = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/);
	// 判断是否为IOS系统
	if (isiOS) {
		setTimeout(() => {
			const scrollHeight = document.documentElement.scrollTop || document.body.scrollTop || 0;
			window.scrollTo(0, Math.max(scrollHeight - 1, 0));
		}, 100);
	}
})


// 生成示例问题
function readQus() {
	var choice_number = 8;
	var request = new XMLHttpRequest();
	request.open("GET", "/statics/dist/examples.json", false);
	request.send(null);
	var factoid_examples = JSON.parse(request.responseText);
	var choiced_factoid_examples = new Array();
	var randomItem = Math.floor(Math.random() * factoid_examples.length);
	var counter = 0
	while (counter < choice_number) {
		choiced_factoid_examples.push(factoid_examples[(randomItem + counter) % factoid_examples.length]);
		counter++;
	}
	for (var i = 0; i < choiced_factoid_examples.length; i++) {
		var p = document.createElement('div');
		p.className = 'ui blue segment example'
		p.innerText = choiced_factoid_examples[i]['query'];
		p.id = 'example_question'
		$('#example_content').append(p);
		//   $('#factoid_modal').append(p);
	}
}

$('#example_content').on('click', ".example", function () {
	$('#example').modal('hide');
	text.val(this.innerText);
	//console.log(text.val());
	action();
	msg_num = msg_num + 1;
	$('.custom.popup').popup('hide');
})

$('body').on('click', ".recommand", function () {
	text.val(this.innerText);
	action();
	msg_num = msg_num + 1;
})

$('#judge').click(function(){
	text.val("病情自测");
	action();
	msg_num = msg_num + 1;
})

$('body').on('click', ".is_button", function () {
	text.val(this.innerText);
	action();
	msg_num = msg_num + 1;
})
