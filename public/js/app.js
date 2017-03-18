var docCookies = {
	getItem: function (sKey) {
		return decodeURIComponent(document.cookie.replace(
			new RegExp("(?:(?:^|.*;)\\s*" +
				encodeURIComponent(sKey).replace(/[\-\.\+\*]/g, "\\$&") +
				"\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
	},
	setItem: function (sKey, sValue, vEnd, sPath, sDomain, bSecure) {
		if (!sKey ||
			/^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)) {
			return false;
		}
		var sExpires = "";
		if (vEnd) {
			switch (vEnd.constructor) {
			case Number:
				sExpires = (vEnd === Infinity) ?
					"; expires=Fri, 31 Dec 9999 23:59:59 GMT" :
					"; max-age=" + vEnd;
				break;
			case String:
				sExpires = "; expires=" + vEnd;
				break;
			case Date:
				sExpires = "; expires=" + vEnd.toGMTString();
				break;
			}
		}
		document.cookie =
			encodeURIComponent(sKey) + "=" +
			encodeURIComponent(sValue) + sExpires +
			(sDomain ? "; domain=" + sDomain : "") +
			(sPath ? "; path=" + sPath : "") +
			(bSecure ? "; secure" : "");

		return true;
	},
	removeItem: function (sKey, sPath) {
		if (!sKey || !this.hasItem(sKey)) {
			return false;
		}
		document.cookie =
			encodeURIComponent(sKey) +
			"=; expires=Thu, 01 Jan 1970 00:00:00 GMT" +
			(sPath ? "; path=" + sPath : "");
		return true;
	},
	hasItem: function (sKey) {
		return (new RegExp("(?:^|;\\s*)" +
			encodeURIComponent(sKey).replace(/[\-\.\+\*]/g, "\\$&") +
			"\\s*\\=")).test(document.cookie);
	},
	keys: /* optional method: you can safely remove it! */ function () {
		var aKeys =
			document.cookie.replace(
				/((?:^|\s*;)[^\=]+)(?=;|$)|^\s*|\s*(?:\=[^;]*)?(?:\1|$)/g,
				"").split(/\s*(?:\=[^;]*)?;\s*/);

		for (var nIdx = 0; nIdx < aKeys.length; nIdx++) {
			aKeys[nIdx] = decodeURIComponent(aKeys[nIdx]);
		}
		return aKeys;
	}
};


docCookies.setItem("test", "none", "", "/", "", "");
var isCookiesEnabled = docCookies.hasItem("test");


if (!isCookiesEnabled) {
	$(".nocookie").show();
	throw new Error("Fuck you nocookie fucklord");
}




function logout () {
	$.get('/users/logout', {},function(e){
			if (e.Code==0){
				alert('OK');
			} else {
				alert('error happened');
			}
			location.reload()
	});
}
var last = 0;
$('#logout-form').hide()
$('#login-form').hide()
var Nickname = '';
$.get('/users/state', {}, function(e){
	
	$('#memory-rss').html(Math.round(e.RSS/1024/1024*1000)/1000+" MiB");
	if (e.Message!= ''){
		$('#logout-form').show()
		$('#name').hide()
	} else {
		$('#login-form').show()
	}
	Nickname = e.Message
	getMessages();
});
$.get('/users/online', {}, function(e){
	var str='';
	var guests=0;
	for (var i = 0; i < e.length; i++) {
		if((e[i]||{}).N == ''){
			guests++;
			continue;
		}
		str += tmpl('online_tmpl', {event: e[i], Nickname:Nickname})
	
	}
	str += tmpl('online_tmpl_guests', {event: guests})
	$('#newMessage #online').html(str);
});

var last_said = 0;
$('#send').click(function(e) {
	/*
	if(last_said + 15 > Date.now()/1000){
		return;
	} else {
		last_said = Date.now()/1000;
	}*/
	
	var message = $('#message').val()
	if(!message.length){
		return;
	}
	$('#message').val('')
	$.post('/messages?user='+ $('#name').val(), {message: message})
});
$('#login').click(function(e) {
	var message = $('#message').val()
	var name = $('#loginname').val()
	var pass = $('#password').val()
	
	$.get('/users/login?name='+name+'&password='+pass, {},function(e){
			if (e.Code==0){
				alert('loged in');
			} else {
				alert('Exit code '+e.Code+'. Error message: '+e.Message);
			}
			location.reload()
	});
	
});
$('#register').click(function(e) {
	var name = $('#loginname').val()
	var pass = $('#password').val()
	
	$.get('/users/register?name='+name+'&password='+pass, {},function(e){
			if (e.Code==0){
				alert('registred');
			} else {
				alert('Exit code '+e.Code+'. Error message: '+e.Message);
			}
			location.reload()
	});
	
});

$('#message').keypress(function(e) {
	if(e.charCode == 13 || e.keyCode == 13) {
		$('#send').click()
		e.preventDefault()
	}
})

var page = 0;
page = parseInt(location.hash.substr(1), 10) || 0;
// Retrieve new messages

if(page == 0){

	$.ajax({
		url: '/page',
		success: function(events) {
			if(events.PrevPage == page){
				return;
			}
			$('#prev-page').html('<a href="#'+events.PrevPage+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + events.PrevPage+'</a>');
		},
		dataType: 'json'
	});

} else {

	$.ajax({
		url: '/page',
		success: function(events) {

			if(page > 1){  
				$('#prev-page').html('<a href="#'+(page-1)+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + (page-1)+'</a>');
			}
			
			if (page < events.PrevPage){ 
				$('#next-page').html('<a href="#'+(page+1)+'" id="next-page" class="page-link" onclick="location.reload()">' + (page+1)+' -></a>');
			} else {
				$('#next-page').html('<a href="#0" id="next-page" class="page-link" onclick="location.reload()">' + (page+1)+' last -></a>');
			}

		},
		dataType: 'json'
	});

}

var getMessages = function() {
	$.ajax({
		url: '/load?page=' +  page,
		success: function(events) {
			$(events).each(function() {
				display(this)
			})
			if(page == 0){
				ketMessages();
			}
			//getMessages()
		},
		dataType: 'json'
	});
}


var ketMessages = function() {
	$.ajax({
		url: '/messages?last=' + last,
		success: function(events) {
			$(events).each(function() {
				display(this)
			})
			ketMessages()
		},
		dataType: 'json'
	});
}

function escapeHtml(html) {
  return String(html)
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#39;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

// Display a message
var display = function(event) {
	last = event.X
	event.M = escapeHtml(event.M)
	event.M = Autolinker.link(event.M);

	$('#thread').append(tmpl('message_tmpl', {event: event, Nickname:Nickname}));
	$('#thread').scrollTo('max')
}


