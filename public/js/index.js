$(document).ready(function() {
  console.log("ready!");
});

document.addEventListener('DOMContentLoaded', function() {
  console.log("ready!");
  var elems = document.querySelectorAll('.modal');
  var instances = M.Modal.init(elems);
});
