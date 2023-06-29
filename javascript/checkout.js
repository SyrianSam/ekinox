$("#submit").click(function(){
    $.post("/calculate", {userinput: $("#userinput").val()}, function(data){
        console.log(data);
        $("#result").text("Total: " + data.totalPrice);
    });
});

$("#generate").click(function() {
    var randomCount = Math.floor(Math.random() * 10) + 1; // Generate a random number between 1 and 10
    var html = "";
  
    for (var i = 1; i <= randomCount; i++) {
      html += "Back to the Future " + i + "\n";
    }
  
    $("#userinput").html(html);
  });