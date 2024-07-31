// Function to enable/disable submit button based on form completeness
function toggleSubmitButton(formId) {
    const form = document.getElementById(formId);
    const submitButton = form.querySelector('.button');
    const inputs = form.querySelectorAll('input[required]');
    let allFilled = true;

    inputs.forEach(input => {
        if (!input.value.trim()) {
            allFilled = false;
        }
    });

    submitButton.disabled = !allFilled;
}

// Ensure DOM is fully loaded before executing scripts
document.addEventListener('DOMContentLoaded', function() {
    // Get elements
    const forgotPasswordLink = document.getElementById('forgot-password-link');
    const backToSignInLink = document.getElementById('back-to-sign-in');
    const forgotPasswordForm = document.getElementById('forgot-password-htm');
    const signInForm = document.querySelector('.sign-in-htm');
    const signUpForm = document.querySelector('.sign-up-htm');

    // Initially hide the forgot password form
    forgotPasswordForm.style.display = 'none';

    // Event listener for showing forgot password form
    forgotPasswordLink.addEventListener('click', (e) => {
        e.preventDefault();
        signInForm.style.display = 'none';
        signUpForm.style.display = 'none';
        forgotPasswordForm.style.display = 'block';
    });

    // Event listener for going back to sign-in form
    backToSignInLink.addEventListener('click', (e) => {
        e.preventDefault();
        forgotPasswordForm.style.display = 'none';
        signInForm.style.display = 'block';
    });

    // Enable/Disable submit button based on form inputs
    document.getElementById('sign-in-form').addEventListener('input', function() {
        toggleSubmitButton('sign-in-form');
    });

    document.getElementById('sign-up-form').addEventListener('input', function() {
        toggleSubmitButton('sign-up-form');
    });

    // Additional setup: Check for initial form state
    toggleSubmitButton('sign-in-form');
    toggleSubmitButton('sign-up-form');
});
