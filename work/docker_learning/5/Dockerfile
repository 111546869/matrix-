# Use the Nginx image from Docker Hub
FROM nginx:latest

# Remove the default Nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf

# Create directories for site A and site B
RUN mkdir -p /usr/share/nginx/html/site_a
RUN mkdir -p /usr/share/nginx/html/site_b

# Copy the custom Nginx config file into the container
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the HTML files to their respective directories
COPY site_a.html /usr/share/nginx/html/site_a/site_a.html
COPY site_b.html /usr/share/nginx/html/site_b/site_b.html

# Expose port 80
EXPOSE 80

# Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]
