var idEvents = {};

$(document).ready(function() {
  $("[id]").each(function(i, elem) {
    var id = elem.id;
    var d = new Date();
    var time = d.getTime()/1000;
    var isInViewport = isElementInViewport(id)
    var isHovered = false;

    idEvents[id] = {
      "clicks": 0,
      "hoverTime": 0,
      "frameTime": 0,
      "visible": isInViewport,
      "becameVisible": isInViewport ? time : -1,
      "enteredHover": -1
    }

    elem = $("#"+id);
    // Handle hover times.
    elem.hover(function() {
      var d = new Date();
      var time = d.getTime()/1000;
      idEvents[id]["enteredHover"] = time;
    }, function() {
      var d = new Date();
      var time = d.getTime()/1000;
      idEvents[id]["hoverTime"] += time - idEvents[id]["enteredHover"];
      console.log("left hover on", id, "after", idEvents[id]["hoverTime"],
                  "sec");
    });
  });

  // Handle scrolls.
  $(window).on('DOMContentLoaded load resize scroll', function() {
    // console.log(idEvents);
    for(id in idEvents) {
      var isVisible = isElementInViewport(id);

      var wasVisible = idEvents[id]["visible"];      
      if(typeof wasVisible === "undefined") {
        wasVisible = isVisible;
        // console.log(id, wasVisible);
      }

      var d = new Date();
      var time = d.getTime()/1000;
      // If the element moved out of view, update its visibility to 'false' and
      // update the time in frame.
      if(wasVisible && !isVisible) {
        idEvents[id]["visible"] = isVisible;
        idEvents[id]["frameTime"] += time - idEvents[id]["becameVisible"];
        console.log(id, "left the view after", idEvents[id]["frameTime"],
                    "sec");
      }
      // If the element entered view, update its visibility to 'true' and reset
      // 'becameVisible'.
      else if(!wasVisible && isVisible) {
        idEvents[id]["visible"] = isVisible;
        idEvents[id]["becameVisible"] = time;
      }
    }
  });

  // Handle element clicks.
  $(document).click(function(e) {
    idEvents[e.target.id]["clicks"] += 1;
    console.log("clicked", e.target.id, idEvents[e.target.id]["clicks"],
                "times");
  });

  // Dump the event data when the user leaves the page.
  $(window).bind("beforeunload", function() {
    $.post("/data", idEvents);
  });
});

function isElementInViewport(id) {
  var elem = $("#"+id)[0];
  var rect = elem.getBoundingClientRect();

  return (
    rect.top >= 0 &&
    rect.left >= 0 &&
    rect.bottom <= (window.innerHeight || $(window).height()) &&
    rect.right <= (window.innerWidth || $(window).width())
  );
}
