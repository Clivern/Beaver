/**
 * Beaver Client
 */

function Socket(url){
    ws = new WebSocket(url);
    ws.onmessage = function(e) { console.log(e); };
    ws.onclose = function(){
        // Try to reconnect in 5 seconds
        setTimeout(function(){Socket(url)}, 5000);
    };
}