$(function(){
    $.getJSON("../assets/json/flickr.json", function(data){
        $.each(data.photosets.photoset, function(key, val){
            var imgTitle = $("#h2" + val.id);
            $(imgTitle).text(val.title._content);
            var imgDescription = $("#p" + val.id);
            $(imgDescription).text(val.description._content);
        });
    });
});