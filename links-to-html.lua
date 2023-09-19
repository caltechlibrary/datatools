-- links-to-html.lua converts links to local Markdown documents to
-- there respective .html counterparts.
function Link(el)
  el.target = string.gsub(el.target, "%.md", ".html")
  return el
end
