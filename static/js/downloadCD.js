(function( window ) {

    var primaryColor;
    var secondaryColor;
    var txtwidth;
    var svgcontent;


    /*function createFavicon(width, height) {
        var width = width;
        var height = height;

        var canvas = document.createElement('canvas');
        canvas.width  = width;
        canvas.height = height;
        var ctx = canvas.getContext('2d');
        var x = canvas.width / 2;
        var y = canvas.height / 2;

        ctx.font = height+"px Startup-Initials";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillStyle = secondaryColor;  
        ctx.fillText($('#SVGID_1_').text() , x, y);

        return canvas.toDataURL("image/png");;
    }*/

    
    
    function createLogoPng(height) {
    	var text_elem = document.getElementById("SVGID_5_");
    	var txtwidth = text_elem.getBBox().width;
    	
    	
    	if (!txtwidth)  ratio = 1;
        txtwidth += 213;
        if (380 > height) ratio = 380/height;
        else ratio = height/380;
    	
    	
    	var svg = $('.svgcontainer').get(0);
        var resizesvg = $($(svg).children().get(0)).clone();
    	
    	var canvas = document.createElement('canvas');
    	//console.log("width: " + (txtwidth*ratio+60));
    	
    	 $(resizesvg).attr({width:(txtwidth*ratio+60), height:height, viewBox:'40 245 400 400'});
        //$(resizesvg).attr({height:height, width:(originwidth*ratio)});

        //$(smallsvg).attr({width:(txtwidth/2.4)+'px', height:myheight+'px', viewBox:'0 245 200 270'});
        //$(smallsvg).attr({transform:scale(0.1,0.1)});
        //$(smallsvg).attr({height:myheight+'px'});
        var svgcontent = $('<div>').append($(resizesvg).clone()).remove().html();
        
        
        canvg(canvas, svgcontent);
       
        cropCanvas(canvas);
      
        var logo = canvas.toDataURL("image/png");
        return logo.substring(22, logo.length);
       
        
    }

    function createTwitterPng() {
        /*var text_elem = document.getElementById("SVGID_5_");
        var txtwidth = text_elem.getBBox().width;
        
        
        if (!txtwidth)  ratio = 1;
        txtwidth += 213;
        if (380 > height) ratio = 380/height;
        else ratio = height/380;*/
       
        var endwidth=300;
        var endheight=300;
        
        var svg = $('.svgcontainer').get(0);
        var resizesvg = $($(svg).children().get(0)).clone();
        
        var canvas = document.createElement('canvas');
        //console.log("width: " + (txtwidth*ratio+60));
        
         $(resizesvg).attr({width:(endwidth*2), height:endheight, viewBox:'40 245 400 400'});
        //$(resizesvg).attr({height:height, width:(originwidth*ratio)});

        //$(smallsvg).attr({width:(txtwidth/2.4)+'px', height:myheight+'px', viewBox:'0 245 200 270'});
        //$(smallsvg).attr({transform:scale(0.1,0.1)});
        //$(smallsvg).attr({height:myheight+'px'});
        var svgcontent = $('<div>').append($(resizesvg).clone()).remove().html();
        
        canvg(canvas, svgcontent);
        
        //console.log("width_alt: " + canvas.width);

        cropCanvas(canvas);

        var ctx = canvas.getContext('2d');
        //var image = ctx.getImageData(0,0,canvas.width,canvas.height);
        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        
        var ratio = endwidth/canvas.width;

        /*console.log("width: " + canvas.width);
        console.log("height: " + canvas.height);*/

        canvasResized.width = canvas.width * ratio;
        canvasResized.height = canvas.height * ratio;
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        canvas.width = 300;
        canvas.height = 300;

        var posY = Math.floor((canvas.height-canvasResized.height) / 2);
        console.log(posY);

        ctx.drawImage(canvasResized, 0, 0, canvasResized.width, canvasResized.height, 0, posY-1, canvasResized.width, canvasResized.height);
        
                

        var logo = canvas.toDataURL("image/png");
        return logo.substring(22, logo.length);

        
    }

    function createCirclePng(height) {
        var width = height;
        
        var svgcont = $('.svgcontainer').get(0);
        var svg = $($(svgcont).children().get(0)).clone();
        var circlesvg = $($('.svg_circles').get(0)).clone();//$('<div>').append($($('.svg_circles').get(0)).clone()).remove().html();
        
        //$(svg).add(circlesvg);
        //var svg = document.createElement('svg');

        var canvas = document.createElement('canvas');

       // $(svg).innerHTML = "";//html(circlesvg);
        $(svg).children().remove();
        $(svg).append(circlesvg);
        $(svg).attr({width:width, height:height});
        
        var svgcontent = $('<div>').append($(svg).clone()).remove().html();
        //console.log(svgcontent);
        
        
        canvg(canvas, svgcontent);
       
        cropCanvas(canvas);
      
        var logo = canvas.toDataURL("image/png");
        return logo.substring(22, logo.length);
       
        
    }

    function createInitialPng(height) {
       
        

        

        var svgcont = $('.svgcontainer').get(0);
        var svg = $($(svgcont).children().get(0)).clone();
        //var initialsvg = $($('.svg_initial').get(0)).clone();//$('<div>').append($($('.svg_circles').get(0)).clone()).remove().html();
        var pattern = $($('defs').get(0)).clone();
        var refinitialsvg = document.getElementById("SVGID_1_");
        var refwidth = refinitialsvg.getBBox().width;
        var refheight = refinitialsvg.getBBox().height;

        var initialsvg = $(refinitialsvg).clone();
        var width = Math.floor((refwidth / refheight) * height)+ 100;

        
        
        /*console.log("refw: " + refwidth);
        console.log("refh: " + refheight);
        console.log("w: " + width);
        console.log("h: " + height);*/

        //$(svg).add(circlesvg);
        //var svg = document.createElement('svg');

        var canvas = document.createElement('canvas');

       // $(svg).innerHTML = "";//html(circlesvg);
        $(svg).children().remove();
        $(svg).append(pattern);
        $(svg).append(initialsvg);
        $(svg).attr({width:width, height:height});
        $(initialsvg).attr({width:width, height:height, transform:"matrix(0.9 0 0 0.9 50 500)" });
        
        var svgcontent = $('<div>').append($(svg).clone()).remove().html();
       // console.log(svgcontent);
       
        
        canvg(canvas, svgcontent);
       
        cropCanvas(canvas);
      
        var logo = canvas.toDataURL("image/png");
        return logo.substring(22, logo.length);
       
        
    }

    function cropCanvas(canvas) {
        var ctx = canvas.getContext('2d');

        var xMin = canvas.width, yMin = canvas.height, xMax = 0, yMax = 0;

        var image = ctx.getImageData(0,0,canvas.width,canvas.height);
        var data = image.data;

        //console.log(data[data.length-1]);

        var hit;
        var gotAyMax = false;
        var mywidth = canvas.width;
        
        //find minimal bounding box
        for (var i = data.length-4; i > 0; i -= 4) {
            
            if (data[i+3] != '0') {
                
                /*if (!gotAyMax) {
                    hit = i / 4;
                    //if (hit==(data.length-4)) hit = hit + 1;
                    gotAyMax = true;
                } */
       
                var tmpYMax = (i/4) / canvas.width;
                if (tmpYMax > yMax) {
                    yMax = tmpYMax;
                }

                var tmpYMin = (i/4) / canvas.width;
                if (tmpYMin < yMin) {
                    yMin = tmpYMin;
                }

                var tmpXMax = (i/4) % canvas.width;
                if (tmpXMax > xMax) {
                    xMax = tmpXMax;
                }

                var tmpXMin = (i/4) % canvas.width;
                if (tmpXMin < xMin) {
                    xMin = tmpXMin;
                }
            }  
        }


        //console.log(hit/4)
        //var yMax = Math.floor((hit)/canvas.width);
        
        
        //console.log("width new: " + xMax);
        //console.log("width old: " + canvas.width);
        
        /*console.log("xMin: " + xMin);
        console.log("xMax: " + xMax);
        console.log("yMin: " + yMin);
        console.log("yMax: " + yMax);*/

        var newimg = ctx.getImageData(xMin,yMin,xMax-xMin+1,yMax-yMin+1);
        
        canvas.height = (yMax-yMin)+2;
        canvas.width = (xMax-xMin)+2;

        ctx.putImageData(newimg, 0, 0);
        /*ctx.strokeStyle = "rgb(255,0,0)";
        ctx.strokeRect(0,0,xMax-xMin,yMax-yMin);*/
    }

    
    $(document).ready(function() {
        var firstclick = false;
        $($('.downloadCD').get(0)).click(function() {
            //prepare and pack all in zip file
            var zip = new JSZip();
            var root = zip.folder("Startup_Live_CD");
            var fonts = root.folder("fonts");
            var logos = root.folder("logos");
            fonts.add("StartupInitials.ttf", "/static/fonts/StartupInitials.ttf");
            fonts.add("StartupText.ttf", "/static/fonts/StartupText.ttf");
            var svg = $($('.svgcontainer').get(0)).html();
            logos.add("svg_logo.svg", svg);
            //console.log(createPng(200));
            logos.add("twitter.png", createTwitterPng(), {base64: true});
            // logos.add("logo_klein.png", createLogoPng(150), {base64: true});
            logos.add("logo_mittel.png", createLogoPng(500), {base64: true});
            // logos.add("logo_gross.png", createLogoPng(1500), {base64: true});
            // logos.add("circles.png", createCirclePng(400), {base64: true});
            // logos.add("initials.png", createInitialPng(400), {base64: true});
            
            /*if (!firstclick) {
                firstclick = true;
            Downloadify.create($('.downloadCD').get(0),{
                swf: '/media/downloadify.swf',
                filename: function(){
                    return "Startup_Live_CD.zip";
                },
                data: function(){ 
                    return zip.generate();
                },
                onComplete: function(){ alert('Your File Has Been Saved!'); },
                onCancel: function(){ alert('You have cancelled the saving of this file.'); },
                onError: function(){ alert('You must put something in the File Contents or there will be nothing to save!'); },
                
                downloadImage: '/images/CDdownload.png',
                width: 100,
                height: 30,
                transparent: true,
                append: true,
                dataType: 'base64'
            });
            }*/
            var content = zip.generate();
            location.href="data:application/zip;base64,"+content;
            
        });


    });

    var logogenSetColors = function(primary, secondary) {
        primaryColor = primary;
        secondaryColor = secondary;
       // alert("pimaryColor set to: " + primaryColor);
    };

    window.logogenSetColors = logogenSetColors;
    
})( window );