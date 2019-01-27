<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width">
    <title>{{.Title}}</title>
    <link href="data:image/png;base64,favicon_will_be_here" rel="icon" type="image/png" />
    <style type="text/css">css_will_be_here</style>
    <style type="text/css" id="theme">theme_will_be_here</style>
    <script>
        const theme = document.getElementById('theme')
        const themeNow = () => localStorage.getItem('theme')
        const setTheme = () => { theme.disabled = themeNow() === 'regular' }
        const toggleTheme = () => localStorage.setItem('theme', (themeNow() === 'regular' ? 'alt' : 'regular')) || setTheme()
        setTheme()
    </script>
    <script>window.onload = function () { js_will_be_here }</script>
</head>

<body>
    <div id="drop-grid"> Drop here to upload </div>
    <input type="file" id="clickupload" style="display:none"/>

    <h1>.{{.Title}}</h1>

    <div id="icHolder">
        <div style="display:none;" class="ic icon-large-images" onclick="window.picsToggle()"></div>
        <div class="ic icon-large-folder" onclick="window.mkdirBtn()"></div>
    </div>

    <div id="pics" style="display:none;">
        <div onclick="window.picsToggle()" id="picsToggleCinema"></div> <img onclick="window.picsNav()" id="picsHolder" />
    </div>

    <table>
    {{range .RowsFolders}}
        <tr>
            <td class="iconRow"><i ondblclick="return rm(event)" onclick="return rename(event)" class="btn icon icon-{{.Ext}} icon-blank"></i></td>
            <td class="file-size"><code>{{.Size}}</code></td>
            <td class="arrow"><div class="arrow-icon"></div></td>
            <td class="display-name"><a class="list-links" onclick="return onClickLink(event)" href="{{.Href}}">{{.Name}}</a></td>
        </tr>
    {{end}}
    {{range .RowsFiles}}
        <tr>
            <td class="iconRow"><i ondblclick="return rm(event)" onclick="return rename(event)" class="btn icon icon-{{.Ext}} icon-blank"></i></td>
            <td class="file-size"><code>{{.Size}}</code></td>
            <td class="arrow"><div class="arrow-icon"></div></td>
            <td class="display-name"><a class="list-links" onclick="return onClickLink(event)" href="{{.Href}}">{{.Name}}</a></td>
        </tr>
    {{end}}
    </table>
</body>
<div id="progress" style="display:none;">
    <span id="dlBarName"></span>
    <div id="dlBarPc">1%</div>
</div>
<div id="ok" class="icon-large-ok"></div>
</html>