function spild() {
    for (var n = 0; n < 10; n++) {
        $("ul").append("<div class='uldiv'> hello world</div>")
    }

}
$(document).ready(function() {
    spild()
    $("li").click(function() {
        $(".uldiv").slideToggle("slow")
    })
})