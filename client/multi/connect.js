var conn;

if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = function (e) {};
    conn.onmessage = function (e) {
        var receiveData = JSON.parse(e.data);
        console.log(receiveData);
    };
    conn.onopen = function (e) {
        conn.send(JSON.stringify(123));
    };
}
