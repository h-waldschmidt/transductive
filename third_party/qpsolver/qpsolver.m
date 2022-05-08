function [x,fval] = qpsolver
H = [1,-1,1
    -1,2,-2
    1,-2,4];
f = [2;-3;1];
lb = zeros(3,1);
ub = ones(size(lb));
Aeq = ones(1,3);
beq = 1/2;
x0 = zeros(3,1);
opts = optimoptions('quadprog','Algorithm','active-set');
[x,fval] = quadprog(H,f,[],[],Aeq,beq,lb,ub,x0,opts)