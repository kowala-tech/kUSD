var src = 'src/';
var dest = 'dist/';

var scripts = [
	'src/js/app.js',
	'src/js/controllers.js',
	'src/js/filters.js',
	'src/js/directives.js',
	'src/js/script.js'
];

var vendor = [
	'dist/js/lib/jquery-1.11.3.min.js',
	'dist/js/lib/angular.min.js',
	'dist/js/lib/ngStorage.min.js',
	'dist/js/lib/lodash.min.js',
	'dist/js/lib/d3.min.js',
	'dist/js/lib/d3.tip.min.js',
	'dist/js/lib/moment.min.js',
	'dist/js/lib/moment.en.min.js',
	'dist/js/lib/toastr.min.js',
	'dist/js/lib/jquery.sparkline.min.js',
	'dist/js/lib/primus.min.js'
];

var styles = [
	'toastr.min.css',
	'style.css'
];

module.exports = function(grunt) {
	grunt.initConfig({
		replace: {
      dist: {
        options: {
          patterns: [
            {
              match: 'cdnURL',
              replacement: '<%= process.env.CDN_URL %>'
            },
						{
              match: 'URL',
              replacement: '<%= process.env.URL %>'
            },
						{
              match: 'explorerURL',
              replacement: '<%= process.env.EXPLORER_URL %>'
            },
						{
              match: 'faucetURL',
              replacement: '<%= process.env.FAUCET_URL %>'
            }
          ]
        },
        files: [
          {expand: true, flatten: true, src: 'dist/index.html', dest: 'dist/'}
        ]
      }
    },
		pkg: grunt.file.readJSON('package.json'),
		clean: {
			build: ['dist'],
			cleanup_js: ['dist/js/*.*', '!dist/js/netstats.*'],
			cleanup_css: ['dist/css/*.css', '!dist/css/netstats.*.css'],
		},
		watch: {
			files: ['src/*/**'],
			tasks: ['build'],
			options: {
				livereload: true
			}
		},
		jade: {
			build: {
				options: {
					data: {
						debug: false,
						pretty: true
					}
				},
				files: {
					'dist/index.html': 'src/views/index.jade',
					'dist/stats/index.html': 'src/views/stats/index.jade'
				}
			}
		},
		copy: {
			build: {
				files: [
					{
						expand: true,
						cwd: 'src/fonts/',
						src: ['minimal-*.*'],
						dest: 'dist/fonts/',
						filter: 'isFile'
					},
					{
						expand: true,
						cwd: 'src/images/',
						src: ['*.ico', '*.png', '*.svg'],
						dest: 'dist/',
						filter: 'isFile'
					},
					{
						expand: true,
						cwd: 'src/css/',
						src: styles,
						dest: 'dist/css/',
						filter: 'isFile'
					},
					{
						expand: true,
						cwd: 'src/js/lib/',
						src: ['*.*'],
						dest: 'dist/js/lib'
					}
				]
			}
		},
		cssmin: {
			build: {
				files: [{
					expand: true,
					cwd: 'dist/css',
					src: ['*.css', '!*.min.css'],
					dest: 'dist/css/'
				}]
			}
		},
		concat: {
			vendor: {
				options: {
					souceMap: false,
					sourceMapIncludeSources: true,
					sourceMapIn: ['dist/js/lib/*.map']
				},
				src: vendor,
				dest: 'dist/js/vendor.min.js'
			},
			scripts : {
				options: {
					separator: ';\n',
				},
				src: scripts,
				dest: 'dist/js/app.js'
			},
			netstats: {
				options: {
					sourceMap: false,
					sourceMapIncludeSources: true,
					sourceMapIn: ['dist/js/vendor.min.js.map', 'dist/js/app.min.js.map']
				},
				src: ['<%= concat.vendor.dest %>', '<%= uglify.app.dest %>'],
				dest: 'dist/js/netstats.min.js',
				nonull: true
			},
			css: {
				src: ['dist/css/*.min.css', 'dist/css/*.css'],
				dest: 'dist/css/netstats.min.css'
			}
		},
		uglify: {
			app: {
				options: {
					mangle: false,
					sourceMap: false,
					sourceMapIncludeSources: true
				},
				dest: 'dist/js/app.min.js',
				src: ['<%= concat.scripts.dest %>']
			}
		}
	});

	grunt.loadNpmTasks('grunt-replace');
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-clean');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-jade');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-uglify');

	grunt.registerTask('default', ['clean:build', 'clean:cleanup_js', 'clean:cleanup_css', 'jade:build', 'copy:build', 'cssmin:build', 'concat:vendor', 'concat:scripts', 'uglify:app', 'concat:netstats', 'concat:css', 'clean:cleanup_js', 'clean:cleanup_css', 'replace']);
	grunt.registerTask('build', 'default');
	grunt.registerTask('all', ['default']);
};
