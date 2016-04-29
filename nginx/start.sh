#!/bin/bash
exec nginx -c /etc/nginx/production.conf -g "daemon off;";
