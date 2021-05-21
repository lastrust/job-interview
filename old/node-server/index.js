const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');
const app = express();

app.use(cors());
app.use(bodyParser.json());
app.post('/test', function(req, res) {
  console.log(req.body);
  res.json(req.body.msg);
});

exports.recruitment = app;
