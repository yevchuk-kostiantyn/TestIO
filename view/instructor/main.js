function getStudentInformation() {
    
}

function getStudentTests() {
    // TODO
}


// Decode JSON and paste it to HTML table
function showResults () {
    var html = '';
    for (var e in json) {
        html += '<tr>'
            +'<td>'+json[e].TestName+'</td>'
            +'<td>'+json[e].Deadline+'</td>'
            +'<td>'+json[e].Grade+'</td>'
        +'</tr>';
    }
    $('#allTests').html(html);
}