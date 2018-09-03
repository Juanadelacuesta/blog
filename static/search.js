$(document).ready(function() {
    $("#q").keyup(function(e) {
        $("#search-results").html('');
        $.get("/search?q=" + e.target.value, function(posts) {
            $.each($.parseJSON(posts), function(i, post) {
                $('<a href="/post/' + post.id + '" class="search-link">' + post.title + '</a>').
                appendTo("#search-results");
            });
        });
    });
});
