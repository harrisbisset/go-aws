function toggleShow(id) {
    let elem = document.getElementById(id)
    console.log(elem.style.visibility)
    if (elem.style.visibility == "hidden") {
        elem.style.visibility = "visible"
    } else {
        elem.style.visibility = "hidden"
    }
}

function toggleShowMore(id) {
    let elem = document.getElementById(id)
    if (elem.innerText == "Show More") {
        elem.innerText = "Hide"
    } else {
        elem.innerText = "Show More"
    }
}
