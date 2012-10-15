/*
 * Photobucket -> Craigslist Code
 *
 * For the P, since she doesn't like copy/pasting
 *
 */

chrome.tabs.getSelected(null, function(tab) {
  var tab_url = tab.url;
  var html = '<center><br />\n  <img src="' + tab.url + '">\n<br /></center>'

  // Add URL
  document.getElementById('url').innerText = html
});
