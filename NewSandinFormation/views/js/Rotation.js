$(function(){

// 设定循环操作

    var time = setInterval(move, 4000);



// 定义图片数组
    var imgArr = ["/views/img/1.jpg", "/views/img/2.jpg", "/views/img/3.jpg", "/views/img/icon.jpg",
        "https://09imgmini.eastday.com/mobile/20200913/2020091319_eb98de66f435478e89219c7b0b7b56c9_4708_mwpm_03200403.jpg"];



// 得到3张图片框

    var leftImg = $("#leftImg");

    var centerImg = $("#centerImg");

    var rightImg = $("#rightImg");



// 设置当前显示的图片

    var currentIndex = 0;



// 设置初始图片

    leftImg.attr("src", imgArr[imgArr.length - 1]);

    centerImg.attr("src", imgArr[0]);

    rightImg.attr("src", imgArr[1]);



// 设置圆点的数量和位置

    for (var i = 0; i < imgArr.length; i++) {

        $("#circleSpan").append("<span></span>");

    }

    $("#circleSpan").css("left", (520 - 20 * imgArr.length)/2 + "px");

    $("#circleSpan span:first").css("background-color", "orange");



// 鼠标悬停和离开事件

    $("#box").hover(function(){

        $("#leftBtn,#rightBtn").show();

// 停止循环

        clearInterval(time);

    }, function(){

        $("#leftBtn,#rightBtn").hide();

// 继续循环

        time = setInterval(move, 4000);

    });



// 给左右按钮添加点击事件

    $("#leftBtn").click(function(){

// 每次滚动显示下一张的图片

        if(currentIndex == 0){

            currentIndex = imgArr.length - 1;

        }else{

            currentIndex--;

        }

// 执行滚动的动画

        leftImg.stop(false, true).animate({"left":"0px"}, 100);

        centerImg.stop(false, true).animate({"left":"520px"}, 100);

        rightImg.stop(false, true).animate({"left": "1040px"}, 102, function(){

            mvoeComplete();

        });

    });

    $("#rightBtn").click(function(){

// 每次滚动显示下一张的图片

        currentIndex++;

// 执行滚动的动画

        leftImg.stop(false, true).animate({"left":"-1040px"}, 100);

        centerImg.stop(false, true).animate({"left":"-520px"}, 100);

        rightImg.stop(false, true).animate({"left": "0px"}, 102, function(){

// 动画结束时执行

            mvoeComplete();

        });

    });



    function move(){

// 每次滚动显示下一张的图片

        currentIndex++;

// 执行滚动的动画

        leftImg.stop(false, true).animate({"left":"-1040px"}, 500);

        centerImg.stop(false, true).animate({"left":"-520px"}, 500);

        rightImg.stop(false, true).animate({"left": "0px"}, 502, function(){

// 动画结束时执行

            mvoeComplete();

        });

    }



    function mvoeComplete(){

// 当滚动结束后位置改变回来

        leftImg.css({"left":"-520px"});

        centerImg.css({"left":"0px"});

        rightImg.css({"left":"520px"});

// 当滚动结束后改变图片显示

        leftImg.attr("src", imgArr[(currentIndex - 1) % imgArr.length]);

        centerImg.attr("src", imgArr[currentIndex % imgArr.length]);

        rightImg.attr("src", imgArr[(currentIndex + 1) % imgArr.length]);



// 改变圆点的颜色

        $("#circleSpan span:eq(" + (currentIndex % imgArr.length) + ")").css("background-color", "orange");

// 将其他的圆点颜色还原成白色

        $("#circleSpan span:not(:eq(" + (currentIndex % imgArr.length) + "))").css("background-color", "white");

    }

});