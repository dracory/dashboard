package adminlte

// templateScript returns the custom JavaScript for the template
func templateScript() string {
	return `
// Wait for both DOM and jQuery to be ready
document.addEventListener('DOMContentLoaded', function() {
    console.log('DOM fully loaded, initializing theme switcher...');
    
    // Function to set a cookie
    function setCookie(name, value, days) {
        try {
            let expires = '';
            if (days) {
                const date = new Date();
                date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
                expires = '; expires=' + date.toUTCString();
            }
            document.cookie = name + '=' + (value || '') + expires + '; path=/';
            console.log('Cookie set:', name, '=', value);
        } catch (e) {
            console.error('Error setting cookie:', e);
        }
    }

    // Function to get a cookie
    function getCookie(name) {
        try {
            const nameEQ = name + '=';
            const ca = document.cookie.split(';');
            for (let i = 0; i < ca.length; i++) {
                let c = ca[i];
                while (c.charAt(0) === ' ') c = c.substring(1, c.length);
                if (c.indexOf(nameEQ) === 0) {
                    const value = c.substring(nameEQ.length, c.length);
                    console.log('Cookie found:', name, '=', value);
                    return value;
                }
            }
            console.log('Cookie not found:', name);
            return null;
        } catch (e) {
            console.error('Error getting cookie:', e);
            return null;
        }
    }

    // Function to apply theme
    function applyTheme(theme) {
        console.log('Applying theme:', theme);
        
        try {
            // Update the theme in the UI
            $('body')
                .removeClass('dark-mode sidebar-dark-primary sidebar-light-primary')
                .removeClass('sidebar-mini sidebar-collapse')
                .addClass('sidebar-mini');
                
            // Apply the selected theme classes
            if (theme === 'dark') {
                $('body').addClass('dark-mode sidebar-dark-primary');
            } else {
                $('body').addClass('sidebar-light-primary');
            }
            
            // Update the active theme indicator
            $('.theme-switch i.fa-check').remove();
            $('.theme-switch[data-theme="' + theme + '"]').append('<i class="fas fa-check float-right mt-1"></i>');
            
            // Save to cookie for persistence
            setCookie('adminlte-theme', theme, 365);
            
            console.log('Theme applied successfully:', theme);
        } catch (e) {
            console.error('Error applying theme:', e);
        }
    }

    // Initialize theme from URL or cookie
    function initTheme() {
        console.log('Initializing theme...');
        
        try {
            // Get theme from URL
            const urlParams = new URLSearchParams(window.location.search);
            const themeParam = urlParams.get('theme');
            
            let theme = 'light'; // Default theme
            
            // Check URL parameter first
            if (themeParam === 'light' || themeParam === 'dark') {
                console.log('Theme from URL parameter:', themeParam);
                theme = themeParam;
            } 
            // Then check cookie
            else {
                const savedTheme = getCookie('adminlte-theme');
                if (savedTheme === 'light' || savedTheme === 'dark') {
                    console.log('Theme from cookie:', savedTheme);
                    theme = savedTheme;
                } else {
                    console.log('Using default theme: light');
                }
            }
            
            // Apply the theme
            applyTheme(theme);
            console.log('Theme initialization complete');
            
        } catch (e) {
            console.error('Error initializing theme:', e);
        }
    }

    // Handle theme switching from the dropdown
    function initThemeSwitcher() {
        console.log('Initializing theme switcher...');
        
        // Use event delegation for dynamically added elements
        $(document).on('click', '.theme-switch', function(e) {
            e.preventDefault();
            const newTheme = $(this).data('theme');
            console.log('Theme switch clicked. New theme:', newTheme);
            
            try {
                // Update URL without page reload
                const url = new URL(window.location);
                url.searchParams.set('theme', newTheme);
                window.history.pushState({}, '', url.toString());
                
                // Apply the new theme
                applyTheme(newTheme);
                
                console.log('Theme switched to:', newTheme);
            } catch (e) {
                console.error('Error switching theme:', e);
            }
            
            return false;
        });
        
        console.log('Theme switcher initialized');
    }

    // Initialize everything when the DOM is ready
    initTheme();
    initThemeSwitcher();
    
    // Debug: Log all theme switch elements
    console.log('Theme switch elements found:', $('.theme-switch').length);
    
    // Debug: Log jQuery version
    console.log('jQuery version:', $.fn.jquery);
});
`
}
