package main

var IndexHTML = `
<html>
  <head>
    <title>Airmeter</title>
    <script src="https://code.jquery.com/jquery-3.3.1.min.js" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    <script defer src="https://use.fontawesome.com/releases/v5.0.6/js/all.js"></script>
  </head>
  
  <body>
      <div class="container-fluid">
        <nav class="navbar navbar-dark bg-dark">
            <a class="navbar-brand" href="#">
              <i class="fas fa-thermometer-empty"></i> Airmeter
            </a>
          </nav>
        <div class="row justify-content-center">
          <div class="col text-center">
            <ul class="list-group" id="reading"></ul>
          </div>
        </div>
      </div>
      <nav class="navbar fixed-bottom navbar-dark bg-dark"></nav>
  </body>

  <script type="text/javascript">
    var endpoint = window.location.protocol + '//' + window.location.host + '/api/reading';
    function getReading(endpoint) {
        (function (endpoint) {
          var startTime = new Date().getTime();
          console.log(startTime + ": Getting information from " + endpoint);
          $.ajax({
            type: "GET",
            url: endpoint,
            timeout: 3000,
            success: function(data){
              total = new Date().getTime()-startTime;
              console.log(name + " start time: " + startTime + " |Total time: "+ total);
              for (var item in data) {
                $('#reading').append('<li class="list-group-item list-group-item-success display-4">' + item + ' : ' + data[item] + ' (' + total + 'ms)</li>')
              }
            },
            error: function(){
              total = new Date().getTime()-startTime;
              console.log("Failed to get information from " + endpoint);
            }
          });
        })(endpoint);
    };

    $(document).ready(function(){
        getReading(endpoint);
    });
  </script>
</html>
`
