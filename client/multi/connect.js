var conn;

if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = function (e) {};
    conn.onmessage = function (e) {
        console.log(e.data);
    };
}
