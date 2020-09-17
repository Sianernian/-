function AutoScroll(obj) {
    $(obj).find("ul:first").animate({
        marginTOp: "-25px"
    }, 1000, function() {
        $(this).css({
            marginTop: "0px"
        }).find("li:first").appendTo(this);
    });
}
$(document).ready(function() {
    setInterval('AutoScroll(".contant")', 2000)
});