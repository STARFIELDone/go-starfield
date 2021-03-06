Pod::Spec.new do |spec|
  spec.name         = 'Gest'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/EvolutionStellarToken/go-EvolutionStellarToken'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS EvolutionStellarToken Client'
  spec.source       = { :git => 'https://github.com/EvolutionStellarToken/go-EvolutionStellarToken.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gest.framework'

	spec.prepare_command = <<-CMD
    curl https://geststore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gest.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
