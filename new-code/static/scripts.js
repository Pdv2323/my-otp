function sendOTP() {
    const email = document.getElementById('emailInput').value;
    if (email) {
        fetch('/send-otp', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email: email })
        }).then(response => {
            if (response.ok) {
                document.getElementById('emailSection').style.display = 'none';
                document.getElementById('otpSection').style.display = 'block';
            } else {
                alert('Failed to send OTP.');
            }
        }).catch(error => {
            console.error('Error:', error);
            alert('Failed to send OTP.');
        });
    } else {
        alert('Please enter a valid email.');
    }
}

function verifyOTP() {
    const email = document.getElementById('emailInput').value;
    const otp = document.getElementById('otpInput').value;
    if (otp) {
        fetch('/verify-otp', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email: email, otp: otp })
        }).then(response => {
            if (response.ok) {
                alert('OTP Verified Successfully!');
            } else {
                alert('Invalid OTP.');
            }
        }).catch(error => {
            console.error('Error:', error);
            alert('Failed to verify OTP.');
        });
    } else {
        alert('Please enter the OTP.');
    }
}
