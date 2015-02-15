(function (window, document, $) {
    "use strict";
    var primaryColor, secondaryColor, txtwidth, txtwidthPadding, txtregion, svgcontent;

    function wrap_rect_around_text(rect_id, text_id, padding) {
        var text_elem, rect;
        text_elem = document.getElementById(text_id);
        txtregion = $(text_elem).text();
        rect = document.getElementById(rect_id);

        if (text_elem && rect) {
            txtwidth = text_elem.getBBox().width;
            if (txtwidth) {
                txtwidthPadding = (txtwidth + padding * 2 + 2);
                rect.setAttribute('width', txtwidthPadding);
            }
        }
    }

    function createPlainFavicon(width, height) {
        var canvas, ctx, x, y;
        canvas = document.createElement('canvas');
        canvas.width  = width;
        canvas.height = height;
        ctx = canvas.getContext('2d');
        x = canvas.width / 2;
        y = canvas.height / 2;

        ctx.font = height + "px StartupInitials";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillStyle = secondaryColor;
        ctx.fillText($('#SVGID_1_').text(), x, y);

        return canvas;
    }

    function createInitialsCanvas(endwidth, endheight, px, py) {
        // console.log("create initials canvas");
        var multi, patternID, pattern, patternwidth, patternheight, pwidth, pheight,
            patternchildren, patternsvg, viewbox, patterncontent, origincanvas,
            canvas, ctx, cpattern, height, canvasResized, resizeContext, ratio;

        endwidth = endwidth * 2;
        endheight = endheight * 2;

        //changing pattern size - multiplikator: trial and error
        multi = 3;

        //patternwidth = endwidth;
        //patternheight = endheight;

        patternID = $('#SVGID_1_').attr("fill");
        patternID = patternID.slice(4, patternID.length - 1);
        pattern = $(patternID).clone();

        switch (patternID) {
        case "#FillPattern1":
            multi = 3.35;
            px = 10;
            py = 10;
            break;
        case "#FillPattern2":
            multi = 2.98;
            px = 30;
            py = 20;
            break;
        case "#FillPattern3":
            multi = 3;
            px = 25;
            py = 30;
            break;
        case "#FillPattern4":
            multi = 2.7;
            px = 8;
            py = 26;
            break;
        case "#FillPattern9":
            multi = 2.5;
            break;
        default:
            multi = 3;
            break;
        }

        patternwidth = endwidth * multi;
        patternheight = endheight * multi;

        pwidth = $(pattern).attr("width");
        pwidth = pwidth.slice(0, pwidth.length - 2);
        pheight = $(pattern).attr("height");
        pheight = pheight.slice(0, pheight.length - 2);
        /*var px = $(pattern).attr("x");
        if (px) px = px.substr(0,px.length-2);
        else px = 0;
        var py = $(pattern).attr("y");
        if (py) py = py.substr(0,py.length-2);
        else py = 0;
        */

        patternchildren = $(pattern).children();
        patternsvg = document.createElementNS("http://www.w3.org/2000/svg", 'svg');
        //set Pattern svg attributes
        $(patternsvg).attr({version: "1.1", xmlns: "http://www.w3.org/2000/svg", "xmlns:xlink": "http://www.w3.org/1999/xlink", preserveAspectRatio: "midX midY"});
        $(patternsvg).attr({width: patternwidth, height: patternheight}); // , "viewBox":"0 0 " + pwidth + " " + pheight
        viewbox = document.createAttribute("viewBox");
        viewbox.nodeValue = "0 0 " + pwidth + " " + pheight;
        //viewbox.nodeValue="0 0 100 100";
        patternsvg.setAttributeNode(viewbox);
        $($(pattern).find("g")).attr({width: pwidth, height: pheight});

        //append pattern to svg 
        $(patternsvg).append(patternchildren);
        //$(pattern).attr({width:endwidth, height:endheight});


        patterncontent = $('<div>').append($(patternsvg).clone()).remove().html();

        //console.log(patterncontent);

        //draw svg
        // $(patternsvg).insertBefore($('canvas').get(0)); 

        origincanvas = document.createElement("canvas");
        //origincanvas.width = endwidth;//pwidth.slice(0,pwidth.length-2);
        //origincanvas.height = endheight;//pheight.slice(0,pheight.length-2);
        canvg(origincanvas, patterncontent);
        //console.log("***** " + origincanvas);
        //
        //console.log($(patternsvg).width());
        //showroomctx.drawSvg(patterncontent, 0,0, endwidth, endheight);
       // console.log(origincanvas.width+" x "+origincanvas.height);

        //Add offset to FillPattern2
        if (patternID == "#FillPattern2") cropCanvas(origincanvas, Math.floor(endheight / 100));
        if (patternID == "#FillPattern9") cropCanvas(origincanvas, Math.floor(endheight / 65));
        else cropCanvas(origincanvas);
        //cropCanvas(origincanvas);
        //console.log(origincanvas.width+" x "+origincanvas.height);
        //return origincanvas;
        //console.log(origincanvas.width);
        canvas = document.createElement("canvas");
        ctx = canvas.getContext('2d');
        cpattern = ctx.createPattern(origincanvas, "repeat");


        height = endheight+(endheight/6);
        canvas.width = endwidth;
        canvas.height = endheight;
        ctx.font = height+"px StartupInitials";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillStyle = cpattern;  
        // offset vars
        //px = 250+px;
        //py = 250+py;
        
        
        // offset
        // ctx.translate(px, py);

        // draw
        ctx.fillText($('#SVGID_1_').text(), canvas.width/2+px,canvas.height/2+py);
        // ctx.fillText($('#SVGID_1_').text(), canvas.width/2,canvas.height/2);
        
        ctx.translate(-px, -py);
        //console.log("width canvas: " + canvas.width);
        //console.log("height canvas: " + canvas.height);


        canvasResized = document.createElement("canvas");
        resizeContext = canvasResized.getContext("2d");
        
        cropCanvas(canvas);

        ratio = (endheight/2)/canvas.height;
       // console.log("width: " + canvas.width);
        //console.log("height: " + canvas.height);

        canvasResized.width = canvas.width*ratio;
        canvasResized.height = endheight/2;

        //console.log("resized width: " + canvasResized.width);
        //console.log("resized height: " + canvasResized.height);
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        

        return canvasResized;
    }

    function createFancyFavicon(endwidth, endheight) {
        var favicon = createInitialsCanvas(endwidth, endheight, 0,0); 
        return favicon;
    }


    function createTestFancyFavicon(endwidth, endheight, px, py) {
        var showroom = initShowRoom();
        showroom.width =endwidth;
        showroom.height = endheight;
        showroom.getContext("2d").drawImage(createInitialsCanvas(endwidth,endheight,px,py),0,0);
        
    }

    function testSvgToPng(height,px,py) {
        var showroom = document.createElement("canvas");//initShowRoom();
        var ratio =  0.9;
        var size = Math.floor(height*ratio);
        var txtsize = Math.floor(height*0.6245);
        
        var initialscanvas = createInitialsCanvas(size*1.5,size,0,0);
        var circlecanvas = createCircleCanvas(size);
        var textcanvas = createTextCanvas(txtsize);
        showroom.height = height;
        showroom.width = 2.5*height;
        //showroom.height = height;
        //console.log("cwidth: " + initialscanvas.width);
        //console.log("height: " + initialscanvas.width);
        var initialsX = Math.floor((height*1.025)-initialscanvas.width/2);
        var initialsY = height-initialscanvas.height;
        var circleX = Math.floor(height/17.3125);
        var txtY = Math.floor(height/4.7);
        showroom.getContext("2d").drawImage(circlecanvas,circleX,0);
        showroom.getContext("2d").drawImage(initialscanvas,initialsX,initialsY); //height/9.09
        showroom.getContext("2d").drawImage(textcanvas,0,txtY);
        
        cropCanvas(showroom);

        return showroom;

    }

    function svgToPng() {
        console.log("svg to png");
        var svg = $($($('.svgcontainer').get(0)).children().get(0)).clone();
    	//var svgcont = $('.svgcontainer').get(0);
        //var svg = $($(svg).children().get(0));
    	//svgcontent = $('.svgcontainer').html();
    	//var canvas = document.createElement('canvas');
    	//canvg(canvas, svgcontent);
    	//$($(svg).children().get(0)).hide();
    	//var logo = Canvas2Image.saveAsPNG(canvas, true);
    	//$(svg).append(logo).show();
/*
        var ratio;
        var myheight = 100;

        if (!txtwidth)  ratio = 1;
        txtwidth += 213;
        if (txtwidth > 390) ratio = 390/txtwidth;
        else ratio = txtwidth/390;

        ratio = myheight/390;
        //alert("textwidth: " + txtwidth + ", ratio: " + ratio);
        $(smallsvg).attr({width:(txtwidth*ratio+60)+'px', height:myheight, viewBox:'40 245 '+(400*ratio)+' 270'});

        //$(smallsvg).attr({width:(txtwidth/2.4)+'px', height:myheight+'px', viewBox:'0 245 200 270'});
        //$(smallsvg).attr({transform:scale(0.1,0.1)});
        //$(smallsvg).attr({height:myheight+'px'});
        svgcontent = $('<div>').append($(smallsvg).clone()).remove().html();
        var smallcanvas = document.createElement('canvas');
        
        canvg(smallcanvas, svgcontent);
       
      
        
        var smalllogo = Canvas2Image.saveAsPNG(smallcanvas, true);
        //$(svg).append(smalllogo).show();

        var logotext = $(smalllogo).attr('src');
        $($('input[type=hidden][name=Png]').get(0)).attr('value', logotext);*/
        //console.log("svg: " + svg);
        $(svg).attr({width:'1000px', height:'200px'});
        var endheight = 100;

        
        var svgcontent = $('<div>').append($(svg).clone()).remove().html();
        //var svgcontent = $('.svgcontainer').html();
        //console.log(svgcontent);
        var canvas = document.createElement('canvas');
        
        canvg(canvas, svgcontent);
        
        //console.log("width_alt: " + canvas.width);

        cropCanvas(canvas);

        var ctx = canvas.getContext('2d');
        //var image = ctx.getImageData(0,0,canvas.width,canvas.height);
        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        
        var ratio = endheight/canvas.height;

        /*console.log("width: " + canvas.width);
        console.log("height: " + canvas.height);*/

        canvasResized.width = canvas.width * ratio;
        canvasResized.height = canvas.height * ratio;
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);
    
        /*console.log("canvas width: " + canvasResized.width);
        console.log("canvas height: " + canvasResized.height);*/

        //console.log("width_new: " + canvas.width);
        
        var publiclogo = Canvas2Image.saveAsPNG(canvasResized, true);
        //$(svg).append(smalllogo).show();

    

        var logotext = $(publiclogo).attr('src');
     //   console.log(logotext);
        //$($('input[type=hidden][name=PublicPng]').get(0)).attr('value', logotext);
        $($('input[type=hidden][name=Png]').get(0)).attr('value', logotext);
        
    }
    
    function testSvgToPublicPng(height,px,py) {
        var showroom = document.createElement("canvas");//initShowRoom();
        var ratio =  0.83;
        var bgratio =  0.91;
        var size = Math.floor(height*ratio);
        var bgsize = Math.floor(height*bgratio);
        var txtsize = Math.floor(height*0.6);
        
        var initialscanvas = createInitialsCanvas(size*1.5,size,0,0);
        var bginitialscanvas = createInitialsBGCanvas(bgsize, 'white');

        var circlecanvas = createCircleCanvas(size);
        var bgcirclecanvas = createCircleBGCanvas(bgsize, "white");

        //var textcanvas = createTextCanvas(txtsize);
        var textcanvas = createTextWithArrowsCanvas(txtsize);

        showroom.height = height;
        showroom.width = 2.5*height;
        //showroom.height = height;
        //console.log("cwidth: " + initialscanvas.width);
        //console.log("height: " + initialscanvas.width);
        var initialsX = Math.floor((height*1.026)-initialscanvas.width/2);
        var initialsY = height-initialscanvas.height-Math.floor(height*0.045);
        
        var circleX = Math.floor(height/18);
        var circleY = Math.floor(height*0.045)

        var bginitialsX = Math.floor((height*1.025)-bginitialscanvas.width/2);
        var bginitialsY = height-bginitialscanvas.height;
        var bgcircleX = Math.floor(height*0.015);
        var txtY = Math.floor(height/4.5);
        
    

        //create Background
        showroom.getContext("2d").drawImage(bginitialscanvas,bginitialsX,bginitialsY);
        showroom.getContext("2d").drawImage(bgcirclecanvas,bgcircleX,0);

        

        showroom.getContext("2d").drawImage(circlecanvas,circleX,circleY);
        showroom.getContext("2d").drawImage(initialscanvas,initialsX,initialsY); //height/9.09

        
        showroom.getContext("2d").drawImage(textcanvas,0,txtY);
        
        cropCanvas(showroom);

        return showroom;
    }
    
    function svgToHoverPng(height) {
        var showroom = document.createElement("canvas");//initShowRoom();
        var ratio =  0.83;
        var bgratio =  0.91;
       
        var txtsize = Math.floor(height*0.6);

        //var textcanvas = createTextCanvas(txtsize);
        var textcanvas = createHoverText(txtsize);

        showroom.height = height;
        showroom.width = 2.5*height;
        //showroom.height = height;
        //console.log("cwidth: " + initialscanvas.width);
        //console.log("height: " + initialscanvas.width);
       
        var txtY = Math.floor(height/4.5);

        
        showroom.getContext("2d").drawImage(textcanvas,0,txtY);
        

        cropCanvas(showroom);

        showroom.width = showroom.width;
        showroom.height = 120;

        showroom.getContext("2d").drawImage(textcanvas,0,txtY);
                
        return showroom;
    }
    
    function svgToPublicPng(endheight) {
        console.log("svg to public png");
        var svg = $($($('.svgcontainer').get(0)).children().get(0)).clone();
        //var smallsvg = $($(svg).children().get(0)).clone();
        //var svgcontent = $('.svgcontainer').html();

        var refelement = $(svg).find('#svg_content');
        //console.log(refelement);

        var circlesvg = $($('.svg_circles').get(0)).clone();
        $(circlesvg).children().attr({stroke:"white", "stroke-width":"20px"});
        //$($($('.svg_circles').get(0)).parent()).prepend(circlesvg);
        $(refelement).prepend(circlesvg);

        var initialsvg = $($('.svg_initial').get(0)).clone();
        $(initialsvg).attr({fill:"white", stroke:"white", "stroke-width":"20px"});
        $(initialsvg).removeAttr("id");
        $(initialsvg).removeAttr("class");
        //$($($('.svg_initial').get(0)).parent().parent()).prepend(initialsvg);
        $(refelement).prepend(initialsvg);

        /*var arrowsbg = $($('.svg_region').get(0)).clone();
        var currwidth = $(arrowsbg).attr("width");
        $(arrowsbg).attr({width:(40), x:210+txtwidth});*/
        

        //FARBE der pfeile: #5a555a , 4 px hoch, 7 breit , dazwischen 2 px, gesamt 38px hoch, 
       
        var arrowgroup = document.createElementNS("http://www.w3.org/2000/svg",'g');
        var arrowdown = document.createElementNS("http://www.w3.org/2000/svg",'polygon');
        var arrowup = document.createElementNS("http://www.w3.org/2000/svg",'polygon');
        //var arrowup = $($('.svg_region').get(0)).clone();
         
        var mywidth = Math.floor(210+txtwidth)+25;
        $(arrowdown).attr({points: (mywidth) +",438 "+(mywidth-12)+",425 "+(mywidth+12)+",425"});
       // $(arrowup).attr({id:"", class:"", fill:"#5a555a", x:500, rx:0, ry:0, transform:"rotate(-45 100 100)", y:400, width:20, height:(7/arrowheight)})
        $(arrowdown).attr({id:"", class:"", fill:"#5a555a"})
       
        $(arrowup).attr({points: (mywidth-12)+",418 "+(mywidth)+",405 "+(mywidth+12)+",418"});
        $(arrowup).attr({id:"", class:"", fill:"#5a555a"})

        var arrowsbg = $($('.svg_region').get(0)).clone();
        $(arrowsbg).attr({id:"", width:50, x:0, y:0})
        var currwidth = $(arrowsbg).attr("width");
        $(arrowsbg).attr({x:210+txtwidth, y:375.822});
        
        $(arrowgroup).append(arrowsbg);
        $(arrowgroup).append(arrowup);
        $(arrowgroup).append(arrowdown);
        
        //$("#svg_content").append(arrowsbg);
        $(refelement).append(arrowgroup);

        //document.getElementById("SVGID_7_").innerHtml = "<rect id='SVGID_6_' x='192.871' y='375.822' rx='11px' ry='11px' width='247.608' height='90.889' fill='#1B171B' />";
        //$('#SVGID_7_').innerHTML = "<rect id='SVGID_6_' x='192.871' y='375.822' rx='11px' ry='11px' width='247.608' height='90.889' fill='#1B171B' />";
        //$($($('.svg_initial').get(0)).parent()).append("<rect id='SVGID_6_' x='192.871' y='375.822' rx='11px' ry='11px' width='247.608' height='90.889' fill='#1B171B' />");        
        //$(circlesvg).attr({stroke:"black", "stroke-width":"20px"});

        //var canvas = document.createElement('canvas');
        
        

        /*var ratio;
        var myheight = 100;

        if (!txtwidth)  ratio = 1;
        txtwidth += 213;
        if (txtwidth > 380) ratio = 380/txtwidth;
        else ratio = txtwidth/380;

        ratio = myheight/380;
        //alert("textwidth: " + txtwidth + ", ratio: " + ratio);
        $(svg).attr({width:(txtwidth*ratio+60)+'px', height:'100px', viewBox:'40 245 '+(400*ratio)+' 270'});

        //$(smallsvg).attr({width:(txtwidth/2.4)+'px', height:myheight+'px', viewBox:'0 245 200 270'});
        //$(smallsvg).attr({transform:scale(0.1,0.1)});
        //$(smallsvg).attr({height:myheight+'px'});*/
        
        $(svg).attr({width:'1000px'});
        
        var svgcontent = $('<div>').append($(svg).clone()).remove().html();
        //console.log(svgcontent);
        var canvas = document.createElement('canvas');
        
        canvg(canvas, svgcontent);
        
        //console.log("width_alt: " + canvas.width);

        cropCanvas(canvas);

        var ctx = canvas.getContext('2d');
        //var image = ctx.getImageData(0,0,canvas.width,canvas.height);
        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        
        var ratio = endheight/canvas.height;

        /*console.log("width: " + canvas.width);
        console.log("height: " + canvas.height);*/

        canvasResized.width = canvas.width * ratio;
        canvasResized.height = canvas.height * ratio;
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);
    
        /*console.log("canvas width: " + canvasResized.width);
        console.log("canvas height: " + canvasResized.height);*/

        //console.log("width_new: " + canvas.width);
        
        var publiclogo = Canvas2Image.saveAsPNG(canvasResized, true);
        //$(svg).append(smalllogo).show();

    

        var logotext = $(publiclogo).attr('src');
     //   console.log(logotext);
        //$($('input[type=hidden][name=PublicPng]').get(0)).attr('value', logotext);
        
        
    }

    function createTwitterPng(height) {
        var width = height;

        var canvas = testSvgToPng(height,0,0);

        var ctx = canvas.getContext('2d');
        //var image = ctx.getImageData(0,0,canvas.width,canvas.height);
        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        

         var ratio = width/canvas.width;

        canvasResized.width = canvas.width * ratio;
        canvasResized.height = canvas.height * ratio;
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        canvas.width = width;
        canvas.height = height;

        var posY = Math.floor((canvas.height-canvasResized.height) / 2);
        console.log(posY);

        ctx.drawImage(canvasResized, 0, 0, canvasResized.width, canvasResized.height, 0, posY-1, canvasResized.width, canvasResized.height);
        

        return canvas;

    }


    function prepareSVG() {
    	
    	return encodeHex(svgcontent);
    	
    }
    
    //Encodes data to Hex(base16) format
    var digitArray = new Array('0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f');
    function toHex(n){
        var result = ''
        var start = true;
        for (var i=32; i>0;){
            i-=4;
            var digit = (n>>i) & 0xf;
            if (!start || digit != 0){
                start = false;
                result += digitArray[digit];
            }
        }
        return (result==''?'0':result);
    }
    function pad(str, len, pad){
        var result = str;
        for (var i=str.length; i<len; i++){
            result = pad + result;
        }
        return result;
    }

    function encodeHex(str){
        var result = "";
        for (var i=0; i<str.length; i++){
            result += pad(toHex(str.charCodeAt(i)&0xff),2,'0');
        }
        return result;
    }
    
    
    


    function cropCanvas(canvas, offset) {
    	if (typeof offset == "undefined") offset = 0;
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
                    yMax = Math.floor(tmpYMax);
                }

                var tmpYMin = (i/4) / canvas.width;
                if (tmpYMin < yMin) {
                    yMin = Math.floor(tmpYMin);
                }

                var tmpXMax = (i/4) % canvas.width;
                if (tmpXMax > xMax) {
                    xMax = Math.floor(tmpXMax);
                }

                var tmpXMin = (i/4) % canvas.width;
                if (tmpXMin < xMin) {
                    xMin = Math.floor(tmpXMin);
                }
            }  
        }
        //console.log(xMin +", "+yMin+", "+xMax+", "+yMax);
        var newimg = ctx.getImageData(xMin,yMin,xMax-xMin+1,yMax-yMin+1);
        
        canvas.height = (yMax-yMin)+1+(offset*2);
        canvas.width = (xMax-xMin)+1+(offset*2);

        //console.log(canvas.width +" x " +canvas.height);

        ctx.putImageData(newimg, (0+offset), (0+offset));
    }

    

    function createInitialsBGCanvas(endheight, color) {
        console.log("create initials bg canvas");
        endheight = endheight*2;
        var endwidth = endheight;
        var linewidth = endheight * 0.09;
        
        var canvas = document.createElement("canvas");
        var ctx = canvas.getContext('2d');
        
        var height = endheight;
        canvas.width = endwidth;
        canvas.height = endheight;
        ctx.font = height+"px StartupInitials";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillStyle = color;  
        ctx.strokeStyle = color;
        ctx.lineWidth = linewidth;
        
        // draw
        ctx.strokeText($('#SVGID_1_').text(), canvas.width/2,canvas.height/2);
        ctx.fillText($('#SVGID_1_').text(), canvas.width/2,canvas.height/2);
        
        //ctx.fillText($('#SVGID_1_').text(), canvas.width/2,canvas.height/2);
        

        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        cropCanvas(canvas);


        var ratio = (endheight/2)/canvas.height;
       // console.log("width: " + canvas.width);
        //console.log("height: " + canvas.height);

        canvasResized.width = canvas.width*ratio;
        canvasResized.height = endheight/2;

        //console.log("resized width: " + canvasResized.width);
        //console.log("resized height: " + canvasResized.height);
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        

        return canvasResized;
    }

    function createCircleCanvas(height) {
        

        var circlesvg = $($('.svg_circles').get(0)).clone();
        var color = $($(circlesvg).children().get(0)).attr("fill");

        var canvas = document.createElement('canvas');
        
        
        //calc values
        //TODO: redefine just for circle
        /*var strokewidthL = Math.floor(height/13.33);
        var strokewidthM = Math.floor(height/13.33);
        var strokewidthS = Math.floor(height/13.33);
        var radiusL = Math.floor(height/(1.19*2));
        var radiusM = Math.floor(height/(1.81*2));
        var radiusS = Math.floor(height/(3.63*2));*/
        var strokewidthL = Math.floor(height/11.375);
        var strokewidthM = Math.floor(height/11.375);
        var strokewidthS = Math.floor(height/11.375);
        var radiusL = Math.floor(height/(2.19));
        var radiusM = Math.floor(height/(3.309));
        var radiusS = Math.floor(height/(6.5));
        
        var x = height/2;
        var y = height/2;

        //get a reference to the canvas
        var ctx = canvas.getContext("2d");
        canvas.height = height;
        canvas.width = 2*height;
        ctx.strokeStyle = color;
        
        //small circle
        ctx.lineWidth = strokewidthS;
        ctx.beginPath();
        ctx.arc(x, y, radiusS, 0, Math.PI*2, true); 
        ctx.closePath();
        ctx.stroke();

        //medium circle
        ctx.lineWidth = strokewidthM;
        ctx.beginPath();
        ctx.arc(x, y, radiusM, 0, Math.PI*2, true); 
        ctx.closePath();
        ctx.stroke();

        //large circle
        ctx.lineWidth = strokewidthL;
        ctx.beginPath();
        ctx.arc(x, y, radiusL, 0, Math.PI*2, true);
        ctx.closePath(); 
        ctx.stroke();
        
        cropCanvas(canvas);

        return canvas; 
        
    }

    function createCircleBGCanvas(height, color) {
        
        var canvas = document.createElement('canvas');
        canvas.width = height;
        canvas.height = height;
        var ctx = canvas.getContext("2d");

        var radius = Math.floor(height/2);
        
        var x = height/2;
        var y = height/2;

        ctx.fillStyle = color;
        ctx.beginPath();
        ctx.arc(x, y, radius, 0, Math.PI*2, true);
        ctx.closePath(); 
        ctx.fill();
        
        //cropCanvas(canvas);

        return canvas; 
        
    }

     function createTextWithArrowsCanvas(height) {
        var viewboxwidth = Math.floor(txtwidthPadding) + 145 + 40;
        var scalewidth = Math.floor(height/(170/viewboxwidth)) +40;
        var arrowX = Math.floor(txtwidthPadding) + 135;

        var mytextsvg = '<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="'+scalewidth+'" height="'+height+'" viewBox="0 0 '+viewboxwidth+' 170" xml:space="preserve">';
        mytextsvg += '<path fill="#1B171B" d="M285.409,56.235c0,6.075-4.926,11-11,11H64.738c-6.075,0-11-4.925-11-11V12.5c0-6.074,4.925-11,11-11h209.671c6.074,0,11,4.926,11,11V56.235L285.409,56.235z"/>';
        mytextsvg += '<path fill="#1B171B" d="M127.712,126.966c0,6.076-4.926,11-11,11H12.5c-6.074,0-11-4.924-11-11V85.781c0-6.074,4.926-11,11-11h104.212c6.074,0,11,4.926,11,11V126.966z"/>';
        mytextsvg += '<g id="SVGID_7_">';
        mytextsvg += '<rect id="SVGID_6_" class="svg_region" x="135.7" y="74.6" rx="11px" ry="11px" width="'+(txtwidthPadding+40)+'" height="90.889" fill="#1B171B" />';
        mytextsvg += '<text id="SVGID_5_" transform="matrix(1 0 0 1 155.9934 142.9768)" fill="#FFFFFF" font-family="StartupText" font-size="70">'+txtregion+'</text>';
        mytextsvg += '</g><g>';
        mytextsvg += '<rect x="9.476" y="90.266" fill="none" width="109.49" height="32.422"/>';
        mytextsvg += '<text id="SVGID_3_" transform="matrix(1 0 0 1 9.4763 119.5051)" fill="#FFFFFF" font-family="StartupText" font-size="40">^</text>';
        mytextsvg += '</g><g>',
        mytextsvg += '<rect x="75.198" y="18.91" fill="none" width="190.808" height="31.915"/>';
        mytextsvg += '<text id="SVGID_4_" transform="matrix(1 0 0 1 75.1985 48.1487)" fill="#FFFFFF" font-family="StartupText" font-size="40">STARTUP</text>';
        mytextsvg += '<polygon transform="matrix(1 0 0 1 ' + arrowX + ' 103)" points="0,15 30,15 15,0" fill="#5a555a"/>';
        mytextsvg += '<polygon transform="matrix(1 0 0 1 ' + arrowX + ' 123)" points="0,0 30,0 15,15" fill="#5a555a"/>';
        mytextsvg += '</g></svg>';


        /*var svgcont = $('.svgcontainer').get(0);
        var svg = $($(svgcont).children().get(0)).clone();
        var textsvg = $($('#svg_content').children(":gt(1)")).clone();//$('<div>').append($($('.svg_circles').get(0)).clone()).remove().html();

        var svgheight = height * 2;
        var svgwidth = svgheight * 3;
        var vBwidth = txtwidthPadding + 15;
        var vBheight = 215;
        var newsvg = document.createElementNS("http://www.w3.org/2000/svg",'svg');
        //set Pattern svg attributes
        $(newsvg).attr({version:"1.1", xmlns:"http://www.w3.org/2000/svg", "xmlns:xlink":"http://www.w3.org/1999/xlink", preserveAspectRatio:"midX midY"});
        $(newsvg).attr({width:svgwidth, height:svgheight}); // , "viewBox":"0 0 " + pwidth + " " + pheight
        var viewbox = document.createAttribute("viewBox");
        viewbox.nodeValue="0 0 " + vBwidth + " " + vBheight;
        newsvg.setAttributeNode(viewbox);
      

        //append pattern to svg 
        $(newsvg).append(textsvg);
        //$(pattern).attr({width:endwidth, height:endheight});
        
        var textsvgcontent = $('<div>').append($(newsvg).clone()).remove().html();*/
        //console.log(mytextsvg);
        var canvas = document.createElement('canvas');
        canvg(canvas, mytextsvg);


        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        cropCanvas(canvas);

        var ratio = 1;
        if (height < canvas.height) ratio = canvas.height/(height);
        else ratio = (height)/canvas.height;
        //console.log("width: " + canvas.width);
        //console.log("height: " + canvas.height);

        canvasResized.width = canvas.width*ratio;
        canvasResized.height = height;

        //console.log("resized width: " + canvasResized.width);
        //console.log("resized height: " + canvasResized.height);
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        

        return canvasResized;
    }

    function createTextCanvas(height) {
        var viewboxwidth = Math.floor(txtwidthPadding + 145);
        var scalewidth = Math.floor(height/(170/viewboxwidth));

        var mytextsvg = '<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="'+scalewidth+'" height="'+height+'" viewBox="0 0 '+viewboxwidth+' 170" xml:space="preserve">';
        mytextsvg += '<path fill="#1B171B" d="M285.409,56.235c0,6.075-4.926,11-11,11H64.738c-6.075,0-11-4.925-11-11V12.5c0-6.074,4.925-11,11-11h209.671c6.074,0,11,4.926,11,11V56.235L285.409,56.235z"/>';
        mytextsvg += '<path fill="#1B171B" d="M127.712,126.966c0,6.076-4.926,11-11,11H12.5c-6.074,0-11-4.924-11-11V85.781c0-6.074,4.926-11,11-11h104.212c6.074,0,11,4.926,11,11V126.966z"/>';
        mytextsvg += '<g id="SVGID_7_">';
        mytextsvg += '<rect id="SVGID_6_" class="svg_region" x="135.7" y="74.6" rx="11px" ry="11px" width="'+txtwidthPadding+'" height="90.889" fill="#1B171B" />';
        mytextsvg += '<text id="SVGID_5_" transform="matrix(1 0 0 1 155.9934 142.9768)" fill="#FFFFFF" font-family="StartupText" font-size="70">'+txtregion+'</text>';
        mytextsvg += '</g><g>';
        mytextsvg += '<rect x="9.476" y="90.266" fill="none" width="109.49" height="32.422"/>';
        mytextsvg += '<text id="SVGID_3_" transform="matrix(1 0 0 1 9.4763 119.5051)" fill="#FFFFFF" font-family="StartupText" font-size="40">^</text>';
        mytextsvg += '</g><g>',
        mytextsvg += '<rect x="75.198" y="18.91" fill="none" width="190.808" height="31.915"/>';
        mytextsvg += '<text id="SVGID_4_" transform="matrix(1 0 0 1 75.1985 48.1487)" fill="#FFFFFF" font-family="StartupText" font-size="40">STARTUP</text>';
        mytextsvg += '</g></svg>';


        /*var svgcont = $('.svgcontainer').get(0);
        var svg = $($(svgcont).children().get(0)).clone();
        var textsvg = $($('#svg_content').children(":gt(1)")).clone();//$('<div>').append($($('.svg_circles').get(0)).clone()).remove().html();

        var svgheight = height * 2;
        var svgwidth = svgheight * 3;
        var vBwidth = txtwidthPadding + 15;
        var vBheight = 215;
        var newsvg = document.createElementNS("http://www.w3.org/2000/svg",'svg');
        //set Pattern svg attributes
        $(newsvg).attr({version:"1.1", xmlns:"http://www.w3.org/2000/svg", "xmlns:xlink":"http://www.w3.org/1999/xlink", preserveAspectRatio:"midX midY"});
        $(newsvg).attr({width:svgwidth, height:svgheight}); // , "viewBox":"0 0 " + pwidth + " " + pheight
        var viewbox = document.createAttribute("viewBox");
        viewbox.nodeValue="0 0 " + vBwidth + " " + vBheight;
        newsvg.setAttributeNode(viewbox);
      

        //append pattern to svg 
        $(newsvg).append(textsvg);
        //$(pattern).attr({width:endwidth, height:endheight});
        
        var textsvgcontent = $('<div>').append($(newsvg).clone()).remove().html();*/
        //console.log(mytextsvg);
        var canvas = document.createElement('canvas');
        canvg(canvas, mytextsvg);


        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        cropCanvas(canvas);

        var ratio = 1;
        if (height < canvas.height) ratio = canvas.height/(height);
        else ratio = (height)/canvas.height;
        //console.log("width: " + canvas.width);
        //console.log("height: " + canvas.height);

        canvasResized.width = canvas.width*ratio;
        canvasResized.height = height;

        //console.log("resized width: " + canvasResized.width);
        //console.log("resized height: " + canvasResized.height);
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        

        return canvasResized;
    }

    function createHoverText(height) {
        var viewboxwidth = Math.floor(txtwidthPadding) + 145 + 40;
        var scalewidth = Math.floor(height/(170/viewboxwidth)) +40;
        var arrowX = Math.floor(txtwidthPadding) + 135;

        var mytextsvg = '<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="'+scalewidth+'" height="'+height+'" viewBox="0 0 '+viewboxwidth+' 170" xml:space="preserve">';
        mytextsvg += '<path fill="#1B171B" d="M285.409,56.235c0,6.075-4.926,11-11,11H64.738c-6.075,0-11-4.925-11-11V12.5c0-6.074,4.925-11,11-11h209.671c6.074,0,11,4.926,11,11V56.235L285.409,56.235z"/>';
        mytextsvg += '<path fill="#1B171B" d="M127.712,126.966c0,6.076-4.926,11-11,11H12.5c-6.074,0-11-4.924-11-11V85.781c0-6.074,4.926-11,11-11h104.212c6.074,0,11,4.926,11,11V126.966z"/>';
        mytextsvg += '<g id="SVGID_7_">';
        mytextsvg += '<rect id="SVGID_6_" class="svg_region" x="135.7" y="74.6" rx="11px" ry="11px" width="'+(txtwidthPadding+40)+'" height="90.889" fill="#1B171B" />';
        mytextsvg += '<text id="SVGID_5_" transform="matrix(1 0 0 1 155.9934 142.9768)" fill="#FFFFFF" font-family="StartupText" font-size="70">'+txtregion+'</text>';
        mytextsvg += '</g><g>';
        mytextsvg += '<rect x="9.476" y="90.266" fill="none" width="109.49" height="32.422"/>';
        mytextsvg += '<text id="SVGID_3_" transform="matrix(1 0 0 1 9.4763 119.5051)" fill="#FFFFFF" font-family="StartupText" font-size="40">^</text>';
        mytextsvg += '</g><g>',
        mytextsvg += '<rect x="75.198" y="18.91" fill="none" width="190.808" height="31.915"/>';
        mytextsvg += '<text id="SVGID_4_" transform="matrix(1 0 0 1 75.1985 48.1487)" fill="#FFFFFF" font-family="StartupText" font-size="40">STARTUP</text>';
        mytextsvg += '<polygon transform="matrix(1 0 0 1 ' + arrowX + ' 103)" points="0,15 30,15 15,0" fill="white"/>';
        mytextsvg += '<polygon transform="matrix(1 0 0 1 ' + arrowX + ' 123)" points="0,0 30,0 15,15" fill="white"/>';
        mytextsvg += '</g></svg>';


        /*var svgcont = $('.svgcontainer').get(0);
        var svg = $($(svgcont).children().get(0)).clone();
        var textsvg = $($('#svg_content').children(":gt(1)")).clone();//$('<div>').append($($('.svg_circles').get(0)).clone()).remove().html();

        var svgheight = height * 2;
        var svgwidth = svgheight * 3;
        var vBwidth = txtwidthPadding + 15;
        var vBheight = 215;
        var newsvg = document.createElementNS("http://www.w3.org/2000/svg",'svg');
        //set Pattern svg attributes
        $(newsvg).attr({version:"1.1", xmlns:"http://www.w3.org/2000/svg", "xmlns:xlink":"http://www.w3.org/1999/xlink", preserveAspectRatio:"midX midY"});
        $(newsvg).attr({width:svgwidth, height:svgheight}); // , "viewBox":"0 0 " + pwidth + " " + pheight
        var viewbox = document.createAttribute("viewBox");
        viewbox.nodeValue="0 0 " + vBwidth + " " + vBheight;
        newsvg.setAttributeNode(viewbox);
      

        //append pattern to svg 
        $(newsvg).append(textsvg);
        //$(pattern).attr({width:endwidth, height:endheight});
        
        var textsvgcontent = $('<div>').append($(newsvg).clone()).remove().html();*/
        //console.log(mytextsvg);
        var canvas = document.createElement('canvas');
        canvg(canvas, mytextsvg);


        var canvasResized = document.createElement("canvas");
        var resizeContext = canvasResized.getContext("2d");
        
        cropCanvas(canvas);

        var ratio = 1;
        if (height < canvas.height) ratio = canvas.height/(height);
        else ratio = (height)/canvas.height;
        //console.log("width: " + canvas.width);
        //console.log("height: " + canvas.height);

        canvasResized.width = canvas.width*ratio;
        canvasResized.height = height;

        //console.log("resized width: " + canvasResized.width);
        //console.log("resized height: " + canvasResized.height);
        
        resizeContext.drawImage(canvas, 0, 0, canvas.width, canvas.height, 0, 0, canvasResized.width, canvasResized.height);

        

        return canvasResized;
    }

    function initShowRoom() {
        
        var showroom;

        if ($("#svgshowroom").length > 0) {
            showroom = $("#svgshowroom");
            showroom.remove();
        } 
        showroom = document.createElement("canvas");
        $(showroom).attr("id", "svgshowroom");
        
        $(showroom).insertAfter($('.svgcontainer').get(0));  
        //$($('.container').get(0)).append(showroom);
        
        //showroom.width = 700;
        //showroom.height = 500;
        return showroom;
    }

    function getPngForZipJs(canvas) {
        var logo = canvas.toDataURL("image/png");
        return logo.substring(22, logo.length);
    }

    $(document).ready(function() {
    	wrap_rect_around_text('SVGID_6_', 'SVGID_5_', 17);
    	svgcontent = $('.svgcontainer').html();
    	//prepareSVG();
        //svgToPng();
        //svgToPublicPng(100);

        //Tests
        //createTestFancyFavicon(129,129,0,0);
    	//svgToHoverPng(120);
        //svgToPng();
        //canvg(showroom, svgcontent);

        // var showroom = initShowRoom();
        // showroom.width = 700;
        // showroom.height = 500;
        // showroom.getContext("2d").drawImage(testSvgToPng(300),0,0);
        // $('svg').hide();

        // var patternID = $('#SVGID_1_').attr("fill");
        // patternID = patternID.slice(4,patternID.length-1);
        // //console.log(patternID);
        // var pattern = $(patternID);
        // var px = pattern.attr("x");
        // if (px) px = px.substr(0,px.length-2);
        // else px = 0;
        // var py = pattern.attr("y");
        // if (py) py = py.substr(0,py.length-2);
        // else py = 0;
        // // console.log("startx: " + px + ", starty: " + py);
        // $(document).keydown(function(e){

        //     //left
        //     if (e.keyCode == 37) { 
        //         px--;
        //         pattern.attr({"x":(px)+"px"});
        //         //return false;
        //     }
        //     //up
        //     if (e.keyCode == 38) {
        //         py--;
        //         pattern.attr({"y":(py)+"px"});
        //         //return false;
        //     }
        //     //right
        //     if (e.keyCode == 39) {
        //         px++;
        //         pattern.attr({"x":(px)+"px"});
        //         //return false;
        //     }
        //     //down
        //     if (e.keyCode == 40) {
        //         py++;
        //         pattern.attr({"y":(py)+"px"});
        //         //return false;
        //     }
        //     //createTestFancyFavicon(129,129, px, py);
        //     //testSvgToPng(400,px,py);
        //     //console.log("tracex: " + px + ", tracey: " + py);
        // });

        /*$("#SVGID_1_").mousedown(function (e) {
            var startx = e.pageX;
            var starty = e.pageY;
            var patternID = $('#SVGID_1_').attr("fill");
            patternID = patternID.slice(4,patternID.length-1);
            //console.log(patternID);
            var pattern = $(patternID);
            var px = pattern.attr("x");
            px = px.substr(0,px.length-2);
            var py = pattern.attr("y");
            py = py.substr(0,py.length-2);
            console.log("startx: " + px + ", starty: " + py);
            $("#SVGID_1_").mousemove(function (e) {
                var currx = e.pageX;
                var curry = e.pageY;
                pattern.attr({"x":(px+(currx-startx))+"px", "y":(py+(curry-starty))+"px"});
                //pattern.attr();
                
                //console.log("x: " + currx + ", y: " + curry);
                var that = this;
                $('body').mouseup(function () {
                    $(that).unbind('mousemove');
                });
            });
            
        });*/
        
        //svgToPng();
        //svgcontent = $('.svgcontainer').html();
        $($('input[type=hidden][name=Svg]').get(0)).attr('value', prepareSVG());

        $($('.saveform').get(0)).submit(function() {
            
            $($('input[type=hidden][name=Png]').get(0)).attr('value', testSvgToPng(100).toDataURL("image/png"));
            $($('input[type=hidden][name=PublicPng]').get(0)).attr('value', testSvgToPublicPng(120).toDataURL("image/png"));
            $($('input[type=hidden][name=HoverPng]').get(0)).attr('value', svgToHoverPng(120).toDataURL("image/png"));
            $($('input[type=hidden][name=Initial60]').get(0)).attr('value', createInitialsBGCanvas(120,60,0,0).toDataURL("image/png"));
            $($('input[type=hidden][name=Initial]').get(0)).attr('value', createInitialsCanvas(150,100,0,0).toDataURL("image/png"));
            $($('input[type=hidden][name=Circle]').get(0)).attr('value', createCircleCanvas(150).toDataURL("image/png"));
            $($('input[type=hidden][name=Favicon16x16]').get(0)).attr('value', createPlainFavicon(16,16).toDataURL("image/png"));
            $($('input[type=hidden][name=Favicon57x57]').get(0)).attr('value', createFancyFavicon(57,57).toDataURL("image/png"));
            $($('input[type=hidden][name=Favicon72x72]').get(0)).attr('value', createFancyFavicon(72,72).toDataURL("image/png"));
            $($('input[type=hidden][name=Favicon114x114]').get(0)).attr('value', createFancyFavicon(114,114).toDataURL("image/png"));
            $($('input[type=hidden][name=Favicon129x129]').get(0)).attr('value', createFancyFavicon(129,129).toDataURL("image/png"));
            $($('input[type=hidden][name=Svg]').get(0)).attr('value', prepareSVG());
            
            
    	});

        $($('.downloadCD').get(0)).click(function() {
            //prepare and pack all in zip file
            console.log("start packing file");
            var zip = new JSZip();
            var root = zip.folder("Startup_Live_CD");
            var fonts = root.folder("fonts");
            var logos = root.folder("logos");
            // fonts.add("StartupInitials.ttf", "/static/fonts/StartupInitials.ttf");
            fonts.add("StartupText.ttf", "/static/fonts/StartupText.ttf");
            var svg = $($('.svgcontainer').get(0)).html();
            logos.add("svg_logo.svg", svg);
            //console.log(createPng(200));
            logos.add("twitter.png", getPngForZipJs(createTwitterPng(300)), {base64: true});
            // logos.add("logo_klein.png", getPngForZipJs(testSvgToPng(150,0,0)), {base64: true});
            logos.add("logo_mittel.png", getPngForZipJs(testSvgToPng(500,0,0)), {base64: true});
            // logos.add("logo_gross.png", getPngForZipJs(testSvgToPng(1500,0,0)), {base64: true});
            // logos.add("circles.png", getPngForZipJs(createCircleCanvas(400)), {base64: true});
            // logos.add("initials.png", getPngForZipJs(createInitialsCanvas(400,400,0,0)), {base64: true});
            // logos.add("text.png", getPngForZipJs(createTextCanvas(400)), {base64: true});
            console.log("generate zip");
            var content = zip.generate();
            console.log("zip generated");
            location.href="data:application/zip;base64,"+content;
            console.log("downloaded");
        });
    });
    var logogenSetColors = function(primary, secondary) {
        primaryColor = primary;
        secondaryColor = secondary;
        //alert("pimaryColor set to: " + primaryColor);
    };

    window.logogenSetColors = logogenSetColors;

    
})(window, document, jQuery);   
