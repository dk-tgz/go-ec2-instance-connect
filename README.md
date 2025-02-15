How to use:
1. Download it to your PATH
2. Run it: 
`go-ec2-instance-connect-ssh --region us-east-2 --instance i-0....... --user ec2-user`

Defaults:
- user: centos
- region: us-east-2


It assumes you have id_rsa.pub and id_rsa inside `$HOME/.ssh/`.

It assumes that you have awscli configured. 
