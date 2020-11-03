// asciinema embedded player
// Modified by MarvinJWendt

(function () {
    function insertAfter(referenceNode, newNode) {
        referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
    }

    function params(container, script) {
        function format(name) {
            var value = script.getAttribute('data-' + name);
            if (value) {
                return name + '=' + value;
            }
        }

        var options = ['size', 'speed', 'autoplay', 'loop', 'theme', 't', 'preload', 'cols', 'rows'];

        return '?' + options.map(format).filter(Boolean).join('&');
    }

    function locationFromString(string) {
        var parser = document.createElement('a');
        parser.href = string;
        return parser;
    }

    function apiHostFromScript(script) {
        var location = locationFromString(script.src);
        return location.protocol + '//' + location.host;
    }

    function insertPlayer(script) {
        console.log("inserting player")
        // do not insert player if there's one already associated with this script
        if (script.dataset.player) {
            return;
        }

        var apiHost = apiHostFromScript(script);

        var asciicastId = script.id.split('-')[1];

        var container = document.createElement('div');
        container.id = "asciicast-container-" + asciicastId;
        container.className = 'asciicast';
        container.style.display = 'block';
        container.style.float = 'none';
        container.style.overflow = 'hidden';
        container.style.padding = 0;
        container.style.margin = '20px 0';

        insertAfter(script, container);

        var iframe = document.createElement('iframe');
        iframe.src = "https://asciinema.org/a/" + asciicastId + '/embed' + params(container, script);
        iframe.id = "asciicast-iframe-" + asciicastId;
        iframe.name = "asciicast-iframe-" + asciicastId;
        iframe.scrolling = "no";
        iframe.setAttribute('allowFullScreen', 'true');
        iframe.style.overflow = "hidden";
        iframe.style.margin = 0;
        iframe.style.border = 0;
        iframe.style.display = "inline-block";
        iframe.style.width = "100%";
        iframe.style.float = "none";
        iframe.style.visibility = "hidden";
        iframe.onload = function () {
            this.style.visibility = 'visible'
        };

        container.appendChild(iframe);

        function receiveSize(e) {
            if (e.origin === apiHost) {
                var name = e.data[0];
                var data = e.data[1];
                var iframeWindow = iframe.contentWindow || iframe;

                if (e.source == iframeWindow && name == 'asciicast:size') {
                    iframe.style.width = '' + data.width + 'px';
                    iframe.style.height = '' + data.height + 'px';
                }
            }
        }

        window.addEventListener("message", receiveSize, false);

        script.dataset.player = container;
    }

    document.addEventListener('DOMContentLoaded', function () {
        setTimeout(() => {
            var scripts = document.querySelectorAll("a[id^='asciicast-']");
            console.log(scripts)
            scripts.forEach(s => insertPlayer(s))
        }, 1000)
    }, false);
})();