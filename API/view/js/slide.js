function AutoScroll(obj) {
    $(obj).find("ul:first").animate({
        marginTOp: "-25px"
    }, 500, function() {
        $(this).css({
            marginTop: "0px"
        }).find("li:first").appendTo(this);
    });
}
$(document).ready(function() {
    setInterval('AutoScroll("#banner")', 1000)
});