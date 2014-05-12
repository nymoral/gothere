function checkLen(obj, min) {
    if (obj.value.length < min) {
        obj.style.borderColor="#FF0000";
    }
    else
        obj.style.borderColor="#8A9B0F";
}

function checkEmail(obj) {
    var re = /[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;

    if (re.test(obj.value)) {
        obj.style.borderColor="#8A9B0F";
    }
    else
        obj.style.borderColor="#FF0000";
}

function checkMatch(obj, min) {
    s = document.getElementById("firstpassword").value;
    if (obj.value === s && obj.value.length >= min) {
        obj.style.borderColor="#8A9B0F";
    }
    else
        obj.style.borderColor="#FF0000";
}

function validateForm() {
    var fn = document.getElementById("first").value;
    var ln = document.getElementById("last").value;
    var em = document.getElementById("email").value;
    var ps = document.getElementById("firstpassword").value;
    var cp = document.getElementById("cpassword").value;

    if (fn.length < 1 || fn.length > 20) {
        alert("Vardas turi būti 1-20 simbolių ilgio.");
        return false
    }
    if (ln.length < 1 || ln.length > 30) {
        alert("Pavardė turi būti 1-30 simbolių ilgio.");
        return false
    }
    var re = /[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/;
    if (! re.test(em)) {
        alert("Netinkamas el. pašto adresas.");
        return false
    }
    if (ps.length < 6) {
        alert("Slaptažodis turi būti bent 6 simbolių ilgio.");
        return false
    }
    if (ps !== cp) {
        alert("Slaptažodžiai nesutampa");
        return false
    }
    return true;
}

function validateFormShort() {
    var ps = document.getElementById("new").value;
    var cp = document.getElementById("repeat").value;

    if (ps.length < 6) {
        alert("Slaptažodis turi būti bent 6 simbolių ilgio.");
        return false
    }
    if (ps !== cp) {
        alert("Slaptažodžiai nesutampa");
        return false
    }
    return true;
}
