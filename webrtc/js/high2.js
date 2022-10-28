$(window).load(function() { cur_call = null; chatting_with = false;
 
$("#conference").show();
 
$("#ext").hide();
 
$("#backbtn").hide();
 
$("#cidname").hide();
 
$("#callbtn").hide();
 
$("#hupbtn").hide();
 
$("#chatwin").hide();
 
$("#chatmsg").hide();
 
$("#chatsend").hide();
 
$("#webcam").hide();
 
$("#video1").hide();
 
$("#login").keyup(function(event) {
 
if (event.keyCode == 13 && !event.shiftKey) {
 
$("#loginbtn").trigger("click");
 
}
 
});
 
$("#ext").keyup(function(event) {
 
if (event.keyCode == 13 && !event.shiftKey) {
 
    $("#callbtn").trigger("click");
 
}
});
 
 
});


/**                         */


function init() { cur_call = null;
 
    chatting_with = false;
     
    var nameHost;
     
    var which_server;
     
    nameHost = window.location.hostname;
     
    which_server = "wss://" + nameHost + ":" + "3384"; console.error("which_server=", which_server);
     
    ua = new SIP.UA({
     
    wsServers: which_server,
     
    uri: $("#login").val() + "@" + nameHost,
     
    password: $("#passwd").val(), userAgentString: 'SIP.js/0.7.7 Sara', traceSip: true,
     
    });
     
     
     
    $("#cidname").keyup(function(event) {
     
    if (event.keyCode == 13 && !event.shiftKey) {
     
    $("#callbtn").trigger("click");
     
    }
     
    });
     
    setupChat();
     
    $(document).keypress(function(event) {
     
    var key = String.fromCharCode(event.keyCode || event.charCode); var i = parseInt(key);
     
    var tag = event.target.tagName.toLowerCase();
     
     
     
    if (tag != 'input') {
     
    if (key === "#" || key === "*" || key === "0" || (i > 0 && i <=9)) {
     
    cur_call.dtmf(key);
     
    }
     
    }
     
    });
     
    }




    function setupChat() {
 
        $("#chatwin").html("");
         
         
         
        $("#chatsend").click(function() {
         
        if (!cur_call && chatting_with) { return;
         
        }
         
        cur_call.message({
         
        to: chatting_with,
         
        body: $("#chatmsg").val(),
         
        from_msg_name: cur_call.params.caller_id_name, from_msg_number: cur_call.params.caller_id_number
         
        });
         
        $("#chatmsg").val("");
         
        });
         
         
         
        $("#chatmsg").keyup(function(event) {
         
        if (event.keyCode == 13 && !event.shiftKey) {
         
        $("#chatsend").trigger("click");
         
        }
         
        });
         
        }