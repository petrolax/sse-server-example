const express = require('express');
const app = express();

const sseServerHost = process.env.SSE_SERVER_HOST;
const clientPort = process.env.CLIENT_PORT;

app.get('/', (req, res) => {
    res.render('client.ejs', { sseServerHost: sseServerHost });
});

app.listen(clientPort, () => {
    console.log('Сервер запущен на порту %d', clientPort);
});