/*
 * Photobucket -> Craigslist Code
 *
 * For the P, since she doesn't like copy/pasting
 *
 */

chrome.tabs.getSelected(null, function(tab) {
  var tab_url = tab.url;
  var html = '\
  <center>\n<br />\n\
    <a href="http://www.esotours.com">\
     \n<img src="' + tab.url + '">\n</a>\
    \n<br />\n\
  </center>'

  // Add URL
  document.getElementById('url').innerText = html
});
